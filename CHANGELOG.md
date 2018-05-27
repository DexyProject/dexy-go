# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [unreleased]

## Added
 - ```depth``` field added to ```/markets``` endpoint.

### Fixed
 - Fixed onOrders query

## [1.0.2] - 2018-05-17

### Changed
 - Added flags to orders to ensure they are not ```UNDER_FUNDED``` when no tokens are present in the vault.

## [1.0.1] - 2018-05-12

### Added
 - Go dep, package management.
 
### Changed
 - Int type as string in json rather than number when reading.
 - Renamed ```cmd/rest``` to ```cmd/api```

## [1.0.0] - 2018-05-02

### Added
 - History endpoint
 - Orders endpoint
 - Orderbook endpoint
 - Markets endpoint
 - Ticks endpoint
 - Cancelled Watcher
 - Trade Watcher
 - Markets worker
 - Tick worker
