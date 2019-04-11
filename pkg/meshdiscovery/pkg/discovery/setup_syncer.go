package discovery

import (
	"context"
	"time"

	"github.com/solo-io/supergloo/pkg/api/clientset"
	"github.com/solo-io/supergloo/pkg/meshdiscovery/pkg/discovery/istio"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
)

func RunDiscoveryEventLoop(ctx context.Context, cs *clientset.Clientset, customErrHandler func(error)) error {
	ctx = contextutils.WithLogger(ctx, "meshdiscovery-event-loop")
	logger := contextutils.LoggerFrom(ctx)

	errHandler := func(err error) {
		if err == nil {
			return
		}
		logger.Errorf("meshdiscovery error: %v", err)
		if customErrHandler != nil {
			customErrHandler(err)
		}
	}

	plugins := configurePlugins()
	meshDicoverySyncer := NewMeshDiscoverySyncer(cs.Input.Mesh, plugins...)

	if err := startEventLoop(ctx, errHandler, cs, meshDicoverySyncer); err != nil {
		return err
	}

	return nil
}

func configurePlugins() []MeshDiscovery {
	plugins := []MeshDiscovery{
		istio.NewIstioMeshDiscovery(),
	}
	return plugins
}

// start the install event loop
func startEventLoop(ctx context.Context, errHandler func(err error), c *clientset.Clientset, syncers v1.DiscoverySyncer) error {
	meshDiscoveryEmitter := v1.NewDiscoveryEmitter(c.Discovery.Pod, c.Input.Mesh)
	meshDiscoveryEventLoop := v1.NewDiscoveryEventLoop(meshDiscoveryEmitter, syncers)

	watchOpts := clients.WatchOpts{
		Ctx:         ctx,
		RefreshRate: time.Second * 1,
	}

	meshDiscoveryErrs, err := meshDiscoveryEventLoop.Run(nil, watchOpts)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case err := <-meshDiscoveryErrs:
				errHandler(err)
			case <-ctx.Done():
			}
		}
	}()
	return nil
}
