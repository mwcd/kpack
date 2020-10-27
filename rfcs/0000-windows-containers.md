## Problem:

kpack does not support building containers that can run on windows

## Outcomes:

### kpack should run on clusters that have windows workers

kpack service pods and build pods for linux images will need to know not to run on the windows workers to work in a mixed-os cluster.

The kpack release configuration should specify to run pods explicitly on linux os.
The kpack controller should schedule linux builds on linux workers.

**Actions to Take**

Use the following node selector for all service pods and build pods for linux images:

```
nodeSelector:
    kubernetes.io/os: linux
```

### kpack should support creating windows builders

This is currently supported in the [pack cli](https://github.com/buildpacks/pack/issues/469)

**Actions to Take**

- kpack-controller will use a new environment variable `WINDOWS_LIFECYCLE_IMAGE` to reference the windows lifecycle image
- When a Builder or ClusterBuilder is created using a windows stack and windows buildpackages, kpack-controller must create a windows builder using the windows lifecycle image.
    - Image os is determined by the image [os property](https://github.com/opencontainers/image-spec/blob/master/config.md#:~:text=os%20string).

Notes:

- Builders with mixed os images must error gracefully
- Builders should be labeled with `os.kpack.io: windows` for os disambiguation independent of image metadata

### kpack should support running builds on windows workers

**Actions to Take**

- kpack-controller will accept new environment variables for windows images needed to run builds. These will have the same functionality as on linux and will be used for running builds on windows workers.
    - `WINDOWS_BUILD_INIT_IMAGE` - a windows build init image
    - `WINDOWS_COMPLETION_IMAGE` - a windows completion image
    - `WINDOWS_REBASE_IMAGE` - a windows rebase image
    - These will be optional to remove the dependency on windows images
- When an Image is created using a windows Builer or ClusterBuilder labeled with `os.kpack.io: windows`, it will create a build with new properties
    - The Build will use `WINDOWS_BUILD_INIT_IMAGE` and `WINDOWS_COMPLETION_IMAGE` for the respective containers.
    - The Build will be labeled with `os.kpack.io: windows` for os disambiguation
 - When a Build is created with the label `os.kpack.io: windows`, it will use the node selector

 ```
 nodeSelector:
     kubernetes.io/os: windows
 ```

## Complexity:

High

## Prior Art:

[pack cli](https://github.com/buildpacks/pack) has support for windows containers

## Risks:

**Windows images can only run and be created on workers with the same OS version**

Windows images can only run on windows nodes with the same OS. ex: Windows Server LTSC 2019 can only run images that are ltsc2019 or equivalent version [docs](https://docs.microsoft.com/en-us/virtualization/windowscontainers/deploy-containers/version-compatibility?tabs=windows-server-2019%2Cwindows-10-20H2#:~:text=Supports%20process%20isolation)

Ergo:

- Each windows image installed with `kpack` (`build-init`, `rebase`, `completion`, `lifecycle`) must be for the same os version and must match the cluster's windows nodes
- Stacks and buildpackages must be for the same os version and must match the cluster's windows nodes
- OCI images created on windows nodes will match the os version of the node

Handling these edge cases is outside of the scope of the rfc, mitigated with docs.

[Supported windows versions](https://kubernetes.io/docs/setup/production-environment/windows/intro-windows-in-kubernetes/#windows-os-version-support)

Note:
- Once they support Windows containers with Hyper-V isolation in Kubernetes, the limitation and compatibility rules will change.

**A cluster could have multiple windows OS versions across different windows nodes**

This would result in some windows builds scheduled on nodes that will not run.

Handling this edge case is outside the scope of the rfc, mitigated with docs.

**If there are no windows nodes, windows builds will not run**

Handling this edge case is outside the scope of the rfc, mitigated with docs.

## Alternatives:

**Handle edge cases around OS compatibility**

Compatibility concerns:
- Install images must match OS version (`build-init`, `rebase`, `completion`, `lifecycle`)
    - This could be surfaced in the kpack controller's logs and windows builds could be blocked
- Buildpackages and stack images must match OS version
    - This could be surfaced in the builder's status
    - This is probably the easiest validation to handle
- All windows images must match the windows nodes
    - This one seems the hardest and probably out of scope

Note:
- Windows images contain a `os.version` property.
    - It is not in the OCI spec but de-facto present on all Microsoft’s base images (which is all of Windows base images) and is therefore inherited to any windows image

```
crane config mcr.microsoft.com/windows/servercore:ltsc2019-amd64  | jq
{
  "architecture": "amd64",
  "os": "windows",
  "os.version": "10.0.17763.1518",
  "created": "2020-10-01T02:26:38.060161+00:00",
...
```