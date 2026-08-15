import 'package:flutter_test/flutter_test.dart';
import 'package:geolocator/geolocator.dart';
import 'package:mobile/database/app_database.dart';
import 'package:mobile/repositories/location_event_repository.dart';
import 'package:mobile/services/location_event_recorder.dart';
import 'package:mobile/services/location_service.dart';
import 'package:sqflite_common_ffi/sqflite_ffi.dart';

class FakeLocationService extends LocationService {
  @override
  Future<Position?> getCurrentLocation() async {
    return Position(
      longitude: 77.2090,
      latitude: 28.6139,
      timestamp: DateTime.fromMillisecondsSinceEpoch(1000),
      accuracy: 5.0,
      altitude: 0.0,
      altitudeAccuracy: 1.0,
      heading: 0.0,
      headingAccuracy: 1.0,
      speed: 0.0,
      speedAccuracy: 1.0,
    );
  }
}

void main() {
  sqfliteFfiInit();
  databaseFactory = databaseFactoryFfi;

  late LocationEventRepository repository;

  setUp(() async {
    await AppDatabase.instance.close();

    repository = LocationEventRepository();

    final db = await AppDatabase.instance.database;

    await db.delete('location_events');
  });

  tearDown(() async {
    await AppDatabase.instance.close();
  });

  test('records location and stores it in SQLite', () async {
    final recorder = LocationEventRecorder(
      locationService: FakeLocationService(),
      repository: repository,
    );

    final event = await recorder.recordLocation(
      userId: 'user-001',
      deviceId: 'device-001',
    );

    expect(event, isNotNull);

    expect(event!.userId, 'user-001');
    expect(event.deviceId, 'device-001');

    expect(event.latitude, 28.6139);
    expect(event.longitude, 77.2090);

    expect(event.sequenceNumber, 1);
    expect(event.synced, false);

    final storedEvents = await repository.getAllEvents();

    expect(storedEvents.length, 1);
    expect(storedEvents.first.eventId, event.eventId);
  });

  test('increments sequence number for multiple locations', () async {
    final recorder = LocationEventRecorder(
      locationService: FakeLocationService(),
      repository: repository,
    );

    final firstEvent = await recorder.recordLocation(
      userId: 'user-001',
      deviceId: 'device-001',
    );

    final secondEvent = await recorder.recordLocation(
      userId: 'user-001',
      deviceId: 'device-001',
    );

    expect(firstEvent!.sequenceNumber, 1);
    expect(secondEvent!.sequenceNumber, 2);

    final storedEvents = await repository.getAllEvents();

    expect(storedEvents.length, 2);
  });
}