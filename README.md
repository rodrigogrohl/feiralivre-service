# FeiraLivre Service

# Architecture
The application follows the Hexagonal Architecture to keep intput/output flows organized and DDD is the main guide to core domain.

# Internal Organization
Simple follows the architecture described above, combined with, the recommended [Go Project Layout](https://github.com/golang-standards/project-layout).  
Highlights to:  
- `presentation`: is the **primary adapter** (REST, gRPC, Prometheus exposing, etc)
- `application`: holds the **core domain** (based on DDD)
- `infrastructure`: resposible for **secondary adapters** (Databases, Mail Relays, Notifications, etc)

