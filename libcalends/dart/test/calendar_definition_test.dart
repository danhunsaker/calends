import 'package:calends/calends.dart';
import 'package:test/test.dart';

import 'fake_calendar.dart';

void main() {
  group('CalendarDefinition interface', () {
    test('listRegistered', () {
      expect(
        CalendarDefinition.listRegistered(),
        unorderedEquals(['Gregorian', 'Jdc', 'Stardate', 'Tai64', 'Unix']),
      );
    });
  });

  group('FakeCalendar class', () {
    test('name', () {
      expect(FakeCalendar().name, equals('fake'));
    });

    test('defaultFormat', () {
      expect(FakeCalendar().defaultFormat, equals('unknown'));
    });

    test('isRegistered', () {
      expect(FakeCalendar().isRegistered(), isFalse);
    });

    test('register/unregister', () {
      final fake = FakeCalendar();
      expect(fake.isRegistered(), isFalse);
      fake.register();
      expect(fake.isRegistered(), isTrue);
      fake.unregister();
      expect(fake.isRegistered(), isFalse);
    });

    test('toInternal', () {
      expect(FakeCalendar().toInternal(0, 'unknown').Seconds, 0);
    });

    test('fromInternal', () {
      expect(
        FakeCalendar().fromInternal(TAI64TimeMethods.fromDouble(0), 'unknown'),
        '40000000000000000000000000000000000000000000000000000000',
      );
    });

    test('offset', () {
      expect(
        FakeCalendar().offset(TAI64TimeMethods.fromDouble(0), 0).Seconds,
        0,
      );
    });

    test('callbacks', () {
      final fake = FakeCalendar();
      expect(fake.isRegistered(), isFalse);

      fake.register();
      expect(fake.isRegistered(), isTrue);

      final test = Calends('0', 'fake', '');
      expect(test.toString(),
          equals('00000000000000000000000000000000000000000000000000000000'));

      expect(test.date('fake', ''),
          equals('00000000000000000000000000000000000000000000000000000000'));

      expect(test.add('1', 'fake').date('tai64', 'tai64naxur'),
          equals('-3FFFFFFFFFFFFFFF0000000000000000000000000000000000000000'));

      fake.unregister();
      expect(fake.isRegistered(), isFalse);
    });
  });
}
