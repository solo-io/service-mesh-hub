# Manually Generate Relay Certificates

Below is an example script to generate each relay certificate manually.
**Note** You need to set the relay agent DNS name to the same cluster name stored in the Gloo Mesh Management Plane.

```sh
RELAY_ROOT_CERT_NAME=relay-root
RELAY_SERVER_CERT_NAME=relay-server-tls
RELAY_SIGNING_CERT_NAME=relay-tls-signing
echo "creating root cert ..."
openssl req -new -newkey rsa:4096 -x509 -sha256 \
        -days 3650 -nodes -out ${RELAY_ROOT_CERT_NAME}.crt -keyout ${RELAY_ROOT_CERT_NAME}.key \
        -subj "/CN=enterprise-networking-ca" \
        -addext "extendedKeyUsage = clientAuth, serverAuth"
echo "creating grpc server tls cert ..."
# server cert
cat > "${RELAY_SERVER_CERT_NAME}.conf" <<EOF
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
[req_distinguished_name]
[ v3_req ]
basicConstraints = CA:FALSE
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth, serverAuth
subjectAltName = @alt_names
[alt_names]
DNS = *.gloo-mesh
EOF
openssl genrsa -out "${RELAY_SERVER_CERT_NAME}.key" 2048
openssl req -new -key "${RELAY_SERVER_CERT_NAME}.key" -out ${RELAY_SERVER_CERT_NAME}.csr -subj "/CN=enterprise-networking-ca" -config "${RELAY_SERVER_CERT_NAME}.conf"
openssl x509 -req \
  -days 3650 \
  -CA ${RELAY_ROOT_CERT_NAME}.crt -CAkey ${RELAY_ROOT_CERT_NAME}.key \
  -set_serial 0 \
  -in ${RELAY_SERVER_CERT_NAME}.csr -out ${RELAY_SERVER_CERT_NAME}.crt \
  -extensions v3_req -extfile "${RELAY_SERVER_CERT_NAME}.conf"
echo "creating identity server signing cert ..."
MGMT_CLUSTER=mgmt-cluster
MGMT_CONTEXT=kind-mgmt-cluster
REMOTE_CLUSTER=remote-cluster
REMOTE_CONTEXT=kind-remote-cluster
# ensure gloo-mesh namespace exists on both mgmt and remote clusters
for context in ${MGMT_CONTEXT} ${REMOTE_CONTEXT}; do
  kubectl --context ${context} create namespace gloo-mesh
done
# create secrets from certs
# Note: ${RELAY_SERVER_CERT_NAME}-secret must match the server Helm value `relayTlsSecret.Name`
kubectl create secret generic ${RELAY_SERVER_CERT_NAME}-secret \
  --from-file=tls.key=${RELAY_SERVER_CERT_NAME}.key \
  --from-file=tls.crt=${RELAY_SERVER_CERT_NAME}.crt \
  --from-file=ca.crt=${RELAY_ROOT_CERT_NAME}.crt \
  --dry-run=client -oyaml | kubectl apply -f- \
   --context ${MGMT_CONTEXT} \
  --namespace gloo-mesh
# Note: ${RELAY_SIGNING_CERT_NAME}-secret must match the server Helm value `signingTlsSecret.Name`
# Use server cert here instead as a temporary workaround
kubectl create secret generic ${RELAY_SIGNING_CERT_NAME}-secret \
  --from-file=tls.key=${RELAY_SERVER_CERT_NAME}.key \
  --from-file=tls.crt=${RELAY_SERVER_CERT_NAME}.crt \
  --from-file=ca.crt=${RELAY_ROOT_CERT_NAME}.crt \
  --dry-run=client -oyaml | kubectl apply -f- \
  --context ${MGMT_CONTEXT} \
  --namespace gloo-mesh
# Note: ${RELAY_ROOT_CERT_NAME}-tls-secret must match the agent Helm value `relay.rootTlsSecret.Name`
for context in ${MGMT_CONTEXT} ${REMOTE_CONTEXT}; do
echo "creating matching root cert for agent in cluster context ${context}..."
  kubectl create secret generic ${RELAY_ROOT_CERT_NAME}-tls-secret \
  --from-file=ca.crt=${RELAY_ROOT_CERT_NAME}.crt \
  --dry-run=client -oyaml | kubectl apply -f- \
   --context ${context} \
  --namespace gloo-mesh
done
RELAY_CLIENT_CERT_NAME=relay-client-tls
# Client certs need to have DNS name from cluster_name
for cluster in ${MGMT_CLUSTER} ${REMOTE_CLUSTER}; do
echo "Creating relay client-certs for cluster  ${cluster}..."
  # The value here is the most important, the management-plane uses the SAN to figure out the indentity
  # So replace the cluster names here if different from the context names.
  echo "[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
[req_distinguished_name]
[ v3_req ]
basicConstraints = CA:FALSE
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth, serverAuth
subjectAltName = @alt_names
[alt_names]
DNS = ${cluster}" > "${RELAY_CLIENT_CERT_NAME}.conf"
  openssl genrsa -out "${RELAY_CLIENT_CERT_NAME}.key" 2048
  openssl req -new -key "${RELAY_CLIENT_CERT_NAME}.key" -out ${RELAY_CLIENT_CERT_NAME}.csr -subj "/CN=enterprise-networking-ca" -config "${RELAY_CLIENT_CERT_NAME}.conf"
  openssl x509 -req \
    -days 3650 \
    -CA ${RELAY_ROOT_CERT_NAME}.crt -CAkey ${RELAY_ROOT_CERT_NAME}.key \
    -set_serial 0 \
    -in ${RELAY_CLIENT_CERT_NAME}.csr -out ${RELAY_CLIENT_CERT_NAME}.crt \
    -extensions v3_req -extfile "${RELAY_CLIENT_CERT_NAME}.conf"
  kubectl create secret generic ${RELAY_CLIENT_CERT_NAME}-secret \
    --from-file=tls.key=${RELAY_CLIENT_CERT_NAME}.key \
    --from-file=tls.crt=${RELAY_CLIENT_CERT_NAME}.crt \
    --from-file=ca.crt=${RELAY_ROOT_CERT_NAME}.crt \
    --dry-run=client -oyaml | kubectl apply -f- \
    --context kind-${cluster}
    --namespace gloo-mesh
done
```