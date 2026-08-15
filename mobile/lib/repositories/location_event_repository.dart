import '../database/app_database.dart';
import '../models/location_event.dart';

class LocationEventRepository {
  final AppDatabase _database = AppDatabase.instance;

  Future<void> insertEvent(LocationEvent event) async {
    final db = await _database.database;

    await db.insert(
      'location_events',
      event.toMap(),
    );
  }

  Future<List<LocationEvent>> getAllEvents() async {
    final db = await _database.database;

    final result = await db.query(
      'location_events',
      orderBy: 'timestamp ASC, sequence_number ASC',
    );

    return result
        .map((map) => LocationEvent.fromMap(map))
        .toList();
  }

  Future<List<LocationEvent>> getPendingEvents() async {
    final db = await _database.database;

    final result = await db.query(
      'location_events',
      where: 'synced = ?',
      whereArgs: [0],
      orderBy: 'timestamp ASC, sequence_number ASC',
    );

    return result
        .map((map) => LocationEvent.fromMap(map))
        .toList();
  }

  Future<void> markAsSynced(String eventId) async {
    final db = await _database.database;

    await db.update(
      'location_events',
      {'synced': 1},
      where: 'event_id = ?',
      whereArgs: [eventId],
    );
  }

  Future<void> deleteAllEvents() async {
    final db = await _database.database;

    await db.delete('location_events');
  }
}