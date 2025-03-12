# GeoDigitalMap-messageRelayService
构建一个 **Go 版本的 FreeTAKServer**，需要考虑 **更完整的架构设计** 和 **全套功能实现**，包括：

- **WebSocket 连接管理**（客户端注册、注销、心跳检测）
- **消息存储与转发**（点对点、点对多点、广播）
- **消息持久化与事务处理**（存储 PostgreSQL，支持消息状态）
- **多中继级联**（支持多个中继服务互通）
- **管理 API**（提供 HTTP API 供管理端使用）
- **TLS/SSL 支持**（保证连接安全）
- **日志与监控**（跟踪 WebSocket 状态）

## **核心功能详细实现**

### **1. WebSocket 连接管理**

每个客户端的 WebSocket 连接都需要维护，支持点对点、点对多点、广播等功能。