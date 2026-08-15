import 'package:flutter_test/flutter_test.dart';
import 'package:mobile/database/app_database.dart';
import 'package:mobile/models/location_event.dart';
import 'package:mobile/repositories/location_event_repository.dart';
import 'package:sqflite_common_ffi/sqflite_ffi.dart';

void main() {
  sqfliteFfiInit();
  databaseFactory = databaseFactoryFfi;

  late LocationEventRepository repository;

  setUp(() async {
    repository = LocationEventRepository();

    await AppDatabase.instance.close();
  });

  tearDown(() async {
    final db = await AppDatabase.instance.database;

    await db.delete('location_events');

    await AppDatabase.instance.close();
  });

  test('stores and retrieves a location event', () async {
    final event = LocationEvent(
      eventId: 'event-001',
      userId: 'user-001',
      latitude: 28.6139,
      longitude: 77.2090,
      deviceId: 'device-001',
      timestamp: DateTime.fromMillisecondsSinceEpoch(1000),
      sequenceNumber: 1,
    );

    await repository.insertEvent(event);

    final events = await repository.getAllEvents();

    expect(events.length, 1);
    expect(events.first.eventId, 'event-001');
    expect(events.first.latitude, 28.6139);
    expect(events.first.synced, false);
  });

  test('returns only pending events', () async {
    final pendingEvent = LocationEvent(
      eventId: 'event-001',
      userId: 'user-001',
      latitude: 28.6139,
      longitude: 77.2090,
      deviceId: 'device-001',
      timestamp: DateTime.fromMillisecondsSinceEpoch(1000),
      sequenceNumber: 1,
    );

    final syncedEvent = LocationEvent(
      eventId: 'event-002',
      userId: 'user-001',
      latitude: 28.6140,
      longitude: 77.2091,
      deviceId: 'device-001',
      timestamp: DateTime.fromMillisecondsSinceEpoch(2000),
      sequenceNumber: 2,
      synced: true,
    );

    await repository.insertEvent(pendingEvent);
    await repository.insertEvent(syncedEvent);

    final pending = await repository.getPendingEvents();

    expect(pending.length, 1);
    expect(pending.first.eventId, 'event-001');
  });

  test('marks an event as synced', () async {
    final event = LocationEvent(
      eventId: 'event-001',
      userId: 'user-001',
      latitude: 28.6139,
      longitude: 77.2090,
      deviceId: 'device-001',
      timestamp: DateTime.fromMillisecondsSinceEpoch(1000),
      sequenceNumber: 1,
    );

    await repository.insertEvent(event);

    await repository.markAsSynced('event-001');

    final pending = await repository.getPendingEvents();

    expect(pending.isEmpty, true);
  });
}