# Change Log

All notable changes to this project will be documented in this file.

## [1.0.0] - 2022-12-09
 
### Added
- Polishing.
- Service connector.

### Changed
- Removed deploy event tracker and moved it to the contract tracker microservice. 

## [1.0.0-rc.1] - 2022-11-19 

### Added
- Contract listener that tracks newly created books from the contract.

### Changed
- Create book flow: when calling a create API request, the backend returns a signature to the frontend that is needed to deploy a book. Simultaneously, the backend adds a basic info about a book and waits for a deploy event from the contract to fully complete all missing information.   


## [1.0.0-rc.0] - 2022-10-21 

### Added
- First stable version with basic functions: adding, updating, deleting, and getting book. 

