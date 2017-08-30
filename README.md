# 运行Go应用的最小Docker环境

- [hello-v1](hello-v1): 最小的 Go 程序, 没有导入任何包, 禁止 cgo (700+KB)
- [hello-v2](hello-v2): 包含 cgo 和 net 包, 开启 cgo 功能, 全部静态链接 (5+MB)

### build环境

为了完美支持cgo的交叉编译, [hello-v2](hello-v2) 采用了在 Docker 中编译目标的方案, 支持在 macOS 等其它系统开发.
