# StreetMarket Service

To **Run All** Tests, Coverage, Database Loader and Application, execute into your terminal:  
`./scripts/run.sh`    

After run, you can use API:

![Running and Using API](./docs/running.gif)


# Architecture
The application follows the Hexagonal Architecture to keep intput/output flows organized and DDD is the main guide to core domain.
  
## Internal Organization
Simple follows the architecture described above, combined with, the recommended [Go Project Layout](https://github.com/golang-standards/project-layout).  
Highlights to:  
- `presentation`: is the **primary adapter** (REST, gRPC, Prometheus exposing, etc)
- `application`: holds the **core domain** (based on DDD)
- `infrastructure`: resposible for **secondary adapters** (Databases, Mail Relays, Notifications, etc)

![Macroview](./docs/service-diagram.drawio.png)

# Next Steps
1. Set database specific schema
2. Expose Prometheus metrics
3. Evaluate GORM tradeoffs
