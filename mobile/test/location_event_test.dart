import 'package:flutter_test/flutter_test.dart';
import 'package:mobile/models/location_event.dart';

void main() {
  test('LocationEvent converts to and from map', () {
    final event = LocationEvent(
      eventId: 'event-001',
      userId: 'user-001',
      deviceId: 'device-001',
      latitude: 28.6139,
      longitude: 77.2090,
      timestamp: DateTime.fromMillisecondsSinceEpoch(1000),
      sequenceNumber: 1,
    );

    final map = event.toMap();
    final restored = LocationEvent.fromMap(map);

    expect(restored.eventId, 'event-001');
    expect(restored.userId, 'user-001');
    expect(restored.deviceId, 'device-001');
    expect(restored.latitude, 28.6139);
    expect(restored.longitude, 77.2090);
    expect(restored.sequenceNumber, 1);
    expect(restored.synced, false);
  });
}