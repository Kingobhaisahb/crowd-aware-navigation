import 'package:uuid/uuid.dart';

import '../models/location_event.dart';
import '../repositories/location_event_repository.dart';
import 'location_service.dart';

class LocationEventRecorder {
  final LocationService _locationService;
  final LocationEventRepository _repository;

  final Uuid _uuid = const Uuid();

  int _sequenceNumber = 0;

  LocationEventRecorder({
    LocationService? locationService,
    LocationEventRepository? repository,
  })  : _locationService =
            locationService ?? LocationService(),
        _repository =
            repository ?? LocationEventRepository();

  Future<LocationEvent?> recordLocation({
    required String userId,
    required String deviceId,
  }) async {
    final position =
        await _locationService.getCurrentLocation();

    if (position == null) {
      return null;
    }

    _sequenceNumber++;

    final event = LocationEvent(
      eventId: _uuid.v4(),
      userId: userId,
      latitude: position.latitude,
      longitude: position.longitude,
      deviceId: deviceId,
      timestamp: DateTime.now(),
      sequenceNumber: _sequenceNumber,
    );

    await _repository.insertEvent(event);

    return event;
  }
}