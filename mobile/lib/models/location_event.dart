class LocationEvent {
  final String eventId;
  final String userId;
  final String deviceId;
  final double latitude;
  final double longitude;
  final DateTime timestamp;
  final int sequenceNumber;
  final bool synced;

  LocationEvent({
    required this.eventId,
    required this.userId,
    required this.deviceId,
    required this.latitude,
    required this.longitude,
    required this.timestamp,
    required this.sequenceNumber,
    this.synced = false,
  });

  Map<String, dynamic> toMap() {
    return {
      'event_id': eventId,
      'user_id': userId,
      'device_id': deviceId,
      'latitude': latitude,
      'longitude': longitude,
      'timestamp': timestamp.millisecondsSinceEpoch,
      'sequence_number': sequenceNumber,
      'synced': synced ? 1 : 0,
    };
  }

  factory LocationEvent.fromMap(Map<String, dynamic> map) {
    return LocationEvent(
      eventId: map['event_id'],
      userId: map['user_id'],
      deviceId: map['device_id'],
      latitude: map['latitude'],
      longitude: map['longitude'],
      timestamp: DateTime.fromMillisecondsSinceEpoch(
        map['timestamp'],
      ),
      sequenceNumber: map['sequence_number'],
      synced: map['synced'] == 1,
    );
  }
}