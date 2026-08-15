class LocationEvent {
  final String eventId;
  final String userId;
  final double latitude;
  final double longitude;
  final String deviceId;
  final DateTime timestamp;
  final int sequenceNumber;
  final bool synced;

  LocationEvent({
    required this.eventId,
    required this.userId,
    required this.latitude,
    required this.longitude,
    required this.deviceId,
    required this.timestamp,
    required this.sequenceNumber,
    this.synced = false,
  });

  Map<String, dynamic> toMap() {
    return {
      'event_id': eventId,
      'user_id': userId,
      'latitude': latitude,
      'longitude': longitude,
      'device_id': deviceId,
      'timestamp': timestamp.millisecondsSinceEpoch,
      'sequence_number': sequenceNumber,
      'synced': synced ? 1 : 0,
    };
  }

  factory LocationEvent.fromMap(Map<String, dynamic> map) {
    return LocationEvent(
      eventId: map['event_id'] as String,
      userId: map['user_id'] as String,
      latitude: (map['latitude'] as num).toDouble(),
      longitude: (map['longitude'] as num).toDouble(),
      deviceId: map['device_id'] as String,
      timestamp: DateTime.fromMillisecondsSinceEpoch(
        map['timestamp'] as int,
      ),
      sequenceNumber: map['sequence_number'] as int,
      synced: map['synced'] == 1,
    );
  }
}