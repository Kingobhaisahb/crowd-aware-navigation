import 'package:flutter/material.dart';

import '../services/location_event_recorder.dart';
import '../repositories/location_event_repository.dart';
import '../services/firebase_location_sync.dart';

class LocationTestScreen extends StatefulWidget {
  const LocationTestScreen({super.key});

  @override
  State<LocationTestScreen> createState() => _LocationTestScreenState();
}

class _LocationTestScreenState extends State<LocationTestScreen> {
  final LocationEventRecorder _recorder = LocationEventRecorder();

  final LocationEventRepository _repository = LocationEventRepository();

  late final FirebaseLocationSync _firebaseSync;

  String _status = 'No location recorded yet.';
  int _localEvents = 0;
  int _pendingEvents = 0;

  @override
  void initState() {
    super.initState();

    _firebaseSync = FirebaseLocationSync(
      repository: _repository,
    );

    _refreshCounts();
  }

  Future<void> _recordLocation() async {
    setState(() {
      _status = 'Getting location...';
    });

    final event = await _recorder.recordLocation(
      userId: 'test-user',
      deviceId: 'test-device',
    );

    if (event == null) {
      setState(() {
        _status = 'Could not get location.';
      });

      return;
    }

    await _refreshCounts();

    setState(() {
      _status =
          'Location recorded:\n'
          'Latitude: ${event.latitude}\n'
          'Longitude: ${event.longitude}\n'
          'Sequence: ${event.sequenceNumber}';
    });
  }

  Future<void> _syncLocations() async {
    setState(() {
      _status = 'Syncing locations...';
    });

    final syncedCount = await _firebaseSync.syncPendingEvents();

    await _refreshCounts();

    setState(() {
      _status = 'Synced $syncedCount location(s) to Firebase.';
    });
  }

  Future<void> _refreshCounts() async {
    final allEvents = await _repository.getAllEvents();

    final pendingEvents = await _repository.getPendingEvents();

    if (!mounted) {
      return;
    }

    setState(() {
      _localEvents = allEvents.length;
      _pendingEvents = pendingEvents.length;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Offline Location Test'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(20),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              _status,
              style: const TextStyle(fontSize: 16),
            ),

            const SizedBox(height: 30),

            Text(
              'Local events: $_localEvents',
              style: const TextStyle(fontSize: 18),
            ),

            const SizedBox(height: 10),

            Text(
              'Pending sync: $_pendingEvents',
              style: const TextStyle(fontSize: 18),
            ),

            const SizedBox(height: 30),

            ElevatedButton(
              onPressed: _recordLocation,
              child: const Text('Record Location'),
            ),

            const SizedBox(height: 15),

            ElevatedButton(
              onPressed: _syncLocations,
              child: const Text('Sync to Firebase'),
            ),
          ],
        ),
      ),
    );
  }
}