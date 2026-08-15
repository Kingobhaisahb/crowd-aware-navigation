import 'package:firebase_database/firebase_database.dart';

import '../repositories/location_event_repository.dart';

class FirebaseLocationSync {
  final LocationEventRepository repository;
  final FirebaseDatabase database;

  FirebaseLocationSync({
    required this.repository,
    FirebaseDatabase? database,
  }) : database = database ?? FirebaseDatabase.instance;

  Future<int> syncPendingEvents() async {
    final pendingEvents = await repository.getPendingEvents();

    int syncedCount = 0;

    for (final event in pendingEvents) {
      try {
        final eventRef = database
            .ref('location_events')
            .child(event.deviceId)
            .child(event.eventId);

        await eventRef.set({
          'eventId': event.eventId,
          'userId': event.userId,
          'deviceId': event.deviceId,
          'latitude': event.latitude,
          'longitude': event.longitude,
          'timestamp': event.timestamp.millisecondsSinceEpoch,
          'sequenceNumber': event.sequenceNumber,
        });

        // Only mark as synced AFTER Firebase successfully accepts it.
        await repository.markAsSynced(event.eventId);

        syncedCount++;
      } catch (e) {
        print('Failed to sync event ${event.eventId}: $e');
      }
    }

    return syncedCount;
  }
}