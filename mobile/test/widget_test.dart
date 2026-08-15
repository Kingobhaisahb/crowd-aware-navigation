import 'package:flutter_test/flutter_test.dart';

import 'package:mobile/main.dart';
import 'package:mobile/screens/location_test_screen.dart';

void main() {
  TestWidgetsFlutterBinding.ensureInitialized();

  testWidgets('Crowd Navigation app loads', (WidgetTester tester) async {
    await tester.pumpWidget(
      const CrowdNavigationApp(),
    );

    await tester.pump();

    expect(
      find.byType(LocationTestScreen),
      findsOneWidget,
    );
  });
}