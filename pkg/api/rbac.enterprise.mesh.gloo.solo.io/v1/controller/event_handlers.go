// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

    rbac_enterprise_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/rbac.enterprise.mesh.gloo.solo.io/v1"

    "github.com/pkg/errors"
    "github.com/solo-io/skv2/pkg/events"
    "sigs.k8s.io/controller-runtime/pkg/manager"
    "sigs.k8s.io/controller-runtime/pkg/predicate"
    "sigs.k8s.io/controller-runtime/pkg/client"
)

// Handle events for the Role Resource
// DEPRECATED: Prefer reconciler pattern.
type RoleEventHandler interface {
    CreateRole(obj *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error
    UpdateRole(old, new *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error
    DeleteRole(obj *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error
    GenericRole(obj *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error
}

type RoleEventHandlerFuncs struct {
    OnCreate  func(obj *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error
    OnUpdate  func(old, new *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error
    OnDelete  func(obj *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error
    OnGeneric func(obj *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error
}

func (f *RoleEventHandlerFuncs) CreateRole(obj *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error {
    if f.OnCreate == nil {
        return nil
    }
    return f.OnCreate(obj)
}

func (f *RoleEventHandlerFuncs) DeleteRole(obj *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error {
    if f.OnDelete == nil {
        return nil
    }
    return f.OnDelete(obj)
}

func (f *RoleEventHandlerFuncs) UpdateRole(objOld, objNew *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error {
    if f.OnUpdate == nil {
        return nil
    }
    return f.OnUpdate(objOld, objNew)
}

func (f *RoleEventHandlerFuncs) GenericRole(obj *rbac_enterprise_mesh_gloo_solo_io_v1.Role) error {
    if f.OnGeneric == nil {
        return nil
    }
    return f.OnGeneric(obj)
}

type RoleEventWatcher interface {
    AddEventHandler(ctx context.Context, h RoleEventHandler, predicates ...predicate.Predicate) error
}

type roleEventWatcher struct {
    watcher events.EventWatcher
}

func NewRoleEventWatcher(name string, mgr manager.Manager) RoleEventWatcher {
    return &roleEventWatcher{
        watcher: events.NewWatcher(name, mgr, &rbac_enterprise_mesh_gloo_solo_io_v1.Role{}),
    }
}

func (c *roleEventWatcher) AddEventHandler(ctx context.Context, h RoleEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRoleHandler{handler: h}
    if err := c.watcher.Watch(ctx, handler, predicates...); err != nil{
        return err
    }
    return nil
}

// genericRoleHandler implements a generic events.EventHandler
type genericRoleHandler struct {
    handler RoleEventHandler
}

func (h genericRoleHandler) Create(object client.Object) error {
    obj, ok := object.(*rbac_enterprise_mesh_gloo_solo_io_v1.Role)
    if !ok {
        return errors.Errorf("internal error: Role handler received event for %T", object)
    }
    return h.handler.CreateRole(obj)
}

func (h genericRoleHandler) Delete(object client.Object) error {
    obj, ok := object.(*rbac_enterprise_mesh_gloo_solo_io_v1.Role)
    if !ok {
        return errors.Errorf("internal error: Role handler received event for %T", object)
    }
    return h.handler.DeleteRole(obj)
}

func (h genericRoleHandler) Update(old, new client.Object) error {
    objOld, ok := old.(*rbac_enterprise_mesh_gloo_solo_io_v1.Role)
    if !ok {
        return errors.Errorf("internal error: Role handler received event for %T", old)
    }
    objNew, ok := new.(*rbac_enterprise_mesh_gloo_solo_io_v1.Role)
    if !ok {
        return errors.Errorf("internal error: Role handler received event for %T", new)
    }
    return h.handler.UpdateRole(objOld, objNew)
}

func (h genericRoleHandler) Generic(object client.Object) error {
    obj, ok := object.(*rbac_enterprise_mesh_gloo_solo_io_v1.Role)
    if !ok {
        return errors.Errorf("internal error: Role handler received event for %T", object)
    }
    return h.handler.GenericRole(obj)
}

// Handle events for the RoleBinding Resource
// DEPRECATED: Prefer reconciler pattern.
type RoleBindingEventHandler interface {
    CreateRoleBinding(obj *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error
    UpdateRoleBinding(old, new *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error
    DeleteRoleBinding(obj *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error
    GenericRoleBinding(obj *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error
}

type RoleBindingEventHandlerFuncs struct {
    OnCreate  func(obj *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error
    OnUpdate  func(old, new *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error
    OnDelete  func(obj *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error
    OnGeneric func(obj *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error
}

func (f *RoleBindingEventHandlerFuncs) CreateRoleBinding(obj *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error {
    if f.OnCreate == nil {
        return nil
    }
    return f.OnCreate(obj)
}

func (f *RoleBindingEventHandlerFuncs) DeleteRoleBinding(obj *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error {
    if f.OnDelete == nil {
        return nil
    }
    return f.OnDelete(obj)
}

func (f *RoleBindingEventHandlerFuncs) UpdateRoleBinding(objOld, objNew *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error {
    if f.OnUpdate == nil {
        return nil
    }
    return f.OnUpdate(objOld, objNew)
}

func (f *RoleBindingEventHandlerFuncs) GenericRoleBinding(obj *rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding) error {
    if f.OnGeneric == nil {
        return nil
    }
    return f.OnGeneric(obj)
}

type RoleBindingEventWatcher interface {
    AddEventHandler(ctx context.Context, h RoleBindingEventHandler, predicates ...predicate.Predicate) error
}

type roleBindingEventWatcher struct {
    watcher events.EventWatcher
}

func NewRoleBindingEventWatcher(name string, mgr manager.Manager) RoleBindingEventWatcher {
    return &roleBindingEventWatcher{
        watcher: events.NewWatcher(name, mgr, &rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding{}),
    }
}

func (c *roleBindingEventWatcher) AddEventHandler(ctx context.Context, h RoleBindingEventHandler, predicates ...predicate.Predicate) error {
	handler := genericRoleBindingHandler{handler: h}
    if err := c.watcher.Watch(ctx, handler, predicates...); err != nil{
        return err
    }
    return nil
}

// genericRoleBindingHandler implements a generic events.EventHandler
type genericRoleBindingHandler struct {
    handler RoleBindingEventHandler
}

func (h genericRoleBindingHandler) Create(object client.Object) error {
    obj, ok := object.(*rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding)
    if !ok {
        return errors.Errorf("internal error: RoleBinding handler received event for %T", object)
    }
    return h.handler.CreateRoleBinding(obj)
}

func (h genericRoleBindingHandler) Delete(object client.Object) error {
    obj, ok := object.(*rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding)
    if !ok {
        return errors.Errorf("internal error: RoleBinding handler received event for %T", object)
    }
    return h.handler.DeleteRoleBinding(obj)
}

func (h genericRoleBindingHandler) Update(old, new client.Object) error {
    objOld, ok := old.(*rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding)
    if !ok {
        return errors.Errorf("internal error: RoleBinding handler received event for %T", old)
    }
    objNew, ok := new.(*rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding)
    if !ok {
        return errors.Errorf("internal error: RoleBinding handler received event for %T", new)
    }
    return h.handler.UpdateRoleBinding(objOld, objNew)
}

func (h genericRoleBindingHandler) Generic(object client.Object) error {
    obj, ok := object.(*rbac_enterprise_mesh_gloo_solo_io_v1.RoleBinding)
    if !ok {
        return errors.Errorf("internal error: RoleBinding handler received event for %T", object)
    }
    return h.handler.GenericRoleBinding(obj)
}
