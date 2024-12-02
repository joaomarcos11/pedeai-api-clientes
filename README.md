# Clientes Pedeai API

This project uses Go with Chi lightweight web framework.

## Testing

#### To generate test coverage:
- `go test -v -coverprofile coverage.out ./...`

#### To get show total test code coverage:
- `go tool cover -func coverage.out`

#### To generate html file with tests code coverage (DO NOT DISPLAY TOTAL COVERAGE):
- `go tool cover -html coverage.out -o cover.html`

#### To generate report to sonarqube code coverage:
- `go test -json > report.out`

### Sonarqube configuration

follow steps in stackoverflow to ignore tests files from report:
https://stackoverflow.com/questions/52962493/how-to-exclude-golang-tests-structs-and-constants-from-counting-against-code-co

### Running the Application

- Docker
- Docker Compose

### Steps

1. Clone this repository
2. Build the dockerfile image
3. Run `docker-compose up` inside deployments folder
4. Application with be server in localhost port 8081

## Hexagonal Architecture

This project follows the principles of Hexagonal Architecture (also known as Ports and Adapters Architecture). The main goal of this architecture is to create loosely coupled application components that can be easily tested and maintained.

### Key Concepts

- **Domain Layer**: Contains the core business logic and domain entities. This layer is independent of any external systems or frameworks.
- **Application Layer**: Contains the application services that orchestrate the business logic. This layer interacts with the domain layer and external systems through ports.
- **Ports**: Interfaces that define the input and output boundaries of the application. Ports are implemented by adapters.
- **Adapters**: Implementations of the ports that interact with external systems (e.g., databases, messaging systems, external APIs).

### Project Structure

- `domain/entities`: Contains the domain entities and business logic.
- `domain/service`: Contains the application services.
- `domain/ports`: Contains the port interfaces.
- `adapters`: Contains the adapter implementations.
- `controllers`: Contains the REST controllers.