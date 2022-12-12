import 'package:calends/calends.dart';
import 'package:test/test.dart';

void main() {
  Calends.initialize();

  group('TAI64 0', () {
    final tai0 = Calends('0', 'tai64', 'decimal');

    test('TAI64 date', () {
      expect(tai0.date('tai64', 'decimal'), equals('0'),
          reason: 'date mismatch');
    });

    test('duration', () {
      expect(tai0.duration(), equals('0'), reason: 'duration mismatch');
    });

    test('TAI64 endDate', () {
      expect(tai0.endDate('tai64', 'decimal'), equals('0'),
          reason: 'endDate mismatch');
    });

    test('add TAI64 1', () {
      final tai1 = tai0.add('1', 'tai64');

      expect(tai1.date('tai64', 'decimal'), equals('1'),
          reason: 'date mismatch');
      expect(tai1.duration(), equals('-1'), reason: 'duration mismatch');
      expect(tai1.endDate('tai64', 'decimal'), equals('0'),
          reason: 'endDate mismatch');
    });

    test('addFromEnd TAI64 1', () {
      final tai1 = tai0.addFromEnd('1', 'tai64');

      expect(tai1.date('tai64', 'decimal'), equals('0'),
          reason: 'date mismatch');
      expect(tai1.duration(), equals('1'), reason: 'duration mismatch');
      expect(tai1.endDate('tai64', 'decimal'), equals('1'),
          reason: 'endDate mismatch');
    });

    test('subtract TAI64 1', () {
      final tai1 = tai0.subtract('1', 'tai64');

      expect(tai1.date('tai64', 'decimal'), equals('-1'),
          reason: 'date mismatch');
      expect(tai1.duration(), equals('1'), reason: 'duration mismatch');
      expect(tai1.endDate('tai64', 'decimal'), equals('0'),
          reason: 'endDate mismatch');
    });

    test('subtractFromEnd TAI64 1', () {
      final tai1 = tai0.subtractFromEnd('1', 'tai64');

      expect(tai1.date('tai64', 'decimal'), equals('0'),
          reason: 'date mismatch');
      expect(tai1.duration(), equals('-1'), reason: 'duration mismatch');
      expect(tai1.endDate('tai64', 'decimal'), equals('-1'),
          reason: 'endDate mismatch');
    });

    test('next TAI64 1', () {
      final tai1 = tai0.next('1', 'tai64');

      expect(tai1.date('tai64', 'decimal'), equals('0'),
          reason: 'date mismatch');
      expect(tai1.duration(), equals('1'), reason: 'duration mismatch');
      expect(tai1.endDate('tai64', 'decimal'), equals('1'),
          reason: 'endDate mismatch');
    });

    test('previous TAI64 1', () {
      final tai1 = tai0.previous('1', 'tai64');

      expect(tai1.date('tai64', 'decimal'), equals('-1'),
          reason: 'date mismatch');
      expect(tai1.duration(), equals('1'), reason: 'duration mismatch');
      expect(tai1.endDate('tai64', 'decimal'), equals('0'),
          reason: 'endDate mismatch');
    });

    test('withDate TAI64 1', () {
      final tai1 = tai0.withDate('1', 'tai64', 'decimal');

      expect(tai1.date('tai64', 'decimal'), equals('1'),
          reason: 'date mismatch');
      expect(tai1.duration(), equals('-1'), reason: 'duration mismatch');
      expect(tai1.endDate('tai64', 'decimal'), equals('0'),
          reason: 'endDate mismatch');
    });

    test('withEndDate TAI64 1', () {
      final tai1 = tai0.withEndDate('1', 'tai64', 'decimal');

      expect(tai1.date('tai64', 'decimal'), equals('0'),
          reason: 'date mismatch');
      expect(tai1.duration(), equals('1'), reason: 'duration mismatch');
      expect(tai1.endDate('tai64', 'decimal'), equals('1'),
          reason: 'endDate mismatch');
    });

    test('withDuration TAI64 1', () {
      final tai1 = tai0.withDuration('1', 'tai64');

      expect(tai1.date('tai64', 'decimal'), equals('0'),
          reason: 'date mismatch');
      expect(tai1.duration(), equals('1'), reason: 'duration mismatch');
      expect(tai1.endDate('tai64', 'decimal'), equals('1'),
          reason: 'endDate mismatch');
    });

    test('withDurationFromEnd TAI64 1', () {
      final tai1 = tai0.withDurationFromEnd('1', 'tai64');

      expect(tai1.date('tai64', 'decimal'), equals('-1'),
          reason: 'date mismatch');
      expect(tai1.duration(), equals('1'), reason: 'duration mismatch');
      expect(tai1.endDate('tai64', 'decimal'), equals('0'),
          reason: 'endDate mismatch');
    });

    test('merge TAI64 1', () {
      final tai1 = Calends('1', 'tai64', 'decimal');
      final tai2 = tai0.merge(tai1);

      expect(tai2.date('tai64', 'decimal'), equals('0'),
          reason: 'date mismatch');
      expect(tai2.duration(), equals('1'), reason: 'duration mismatch');
      expect(tai2.endDate('tai64', 'decimal'), equals('1'),
          reason: 'endDate mismatch');
    });

    test('intersect TAI64 -1:1', () {
      final tai1 = Calends(
          <String, String>{'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      final tai2 = tai0.intersect(tai1);

      expect(tai2.date('tai64', 'decimal'), equals('0'),
          reason: 'date mismatch');
      expect(tai2.duration(), equals('0'), reason: 'duration mismatch');
      expect(tai2.endDate('tai64', 'decimal'), equals('0'),
          reason: 'endDate mismatch');
    });

    test('gap TAI64 1', () {
      final tai1 = Calends('1', 'tai64', 'decimal');
      final tai2 = tai0.gap(tai1);

      expect(tai2.date('tai64', 'decimal'), equals('0'),
          reason: 'date mismatch');
      expect(tai2.duration(), equals('1'), reason: 'duration mismatch');
      expect(tai2.endDate('tai64', 'decimal'), equals('1'),
          reason: 'endDate mismatch');
    });

    test('difference TAI64 1', () {
      final tai1 = Calends('1', 'tai64', 'decimal');
      expect(tai0.difference(tai1, 'start'), equals('-1'));
    });

    test('compare TAI64 1', () {
      final tai1 = Calends('1', 'tai64', 'decimal');
      expect(tai0.compare(tai1, 'start'), equals(-1));
    });

    test('isSame TAI64 1', () {
      final tai1 = Calends('1', 'tai64', 'decimal');
      expect(tai0.isSame(tai1), isFalse);
    });

    test('contains TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.contains(tai1), isFalse);
    });

    test('overlaps TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.overlaps(tai1), isTrue);
    });

    test('abuts TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.abuts(tai1), isFalse);
    });

    test('isBefore TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.isBefore(tai1), isFalse);
    });

    test('startsBefore TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.startsBefore(tai1), isFalse);
    });

    test('endsBefore TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.endsBefore(tai1), isTrue);
    });

    test('isDuring TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.isDuring(tai1), isTrue);
    });

    test('startsDuring TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.startsDuring(tai1), isTrue);
    });

    test('endsDuring TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.endsDuring(tai1), isTrue);
    });

    test('isAfter TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.isAfter(tai1), isFalse);
    });

    test('startsAfter TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.startsAfter(tai1), isTrue);
    });

    test('endsAfter TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.endsAfter(tai1), isFalse);
    });

    test('isShorter TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.isShorter(tai1), isTrue);
    });

    test('isSameDuration TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.isSameDuration(tai1), isFalse);
    });

    test('isLonger TAI64 -1:1', () {
      final tai1 = Calends({'start': '-1', 'end': '1'}, 'tai64', 'decimal');
      expect(tai0.isLonger(tai1), isFalse);
    });

    test('toString', () {
      expect(tai0.toString(),
          equals('40000000000000000000000000000000000000000000000000000000'));
    });

    test('encodeText', () {
      expect(tai0.encodeText(),
          equals('40000000000000000000000000000000000000000000000000000000'));
    });

    test('encodeJson', () {
      expect(tai0.encodeJson(),
          equals('"40000000000000000000000000000000000000000000000000000000"'));
    });

    test('== TAI64 1', () {
      final tai1 = Calends('1', 'tai64', 'decimal');
      expect(tai0 == tai1, isFalse);
    });

    test('< TAI64 1', () {
      final tai1 = Calends('1', 'tai64', 'decimal');
      expect(tai0 < tai1, isTrue);
    });

    test('<= TAI64 1', () {
      final tai1 = Calends('1', 'tai64', 'decimal');
      expect(tai0 <= tai1, isTrue);
    });

    test('> TAI64 1', () {
      final tai1 = Calends('1', 'tai64', 'decimal');
      expect(tai0 > tai1, isFalse);
    });

    test('>= TAI64 1', () {
      final tai1 = Calends('1', 'tai64', 'decimal');
      expect(tai0 >= tai1, isFalse);
    });
  });

  group('decodeText', () {
    final tai0 = Calends.decodeText(
        '40000000000000000000000000000000000000000000000000000000');

    test('TAI64 date', () {
      expect(tai0.date('tai64', 'decimal'), equals('0'),
          reason: 'date mismatch');
    });

    test('duration', () {
      expect(tai0.duration(), equals('0'), reason: 'duration mismatch');
    });

    test('TAI64 endDate', () {
      expect(tai0.endDate('tai64', 'decimal'), equals('0'),
          reason: 'endDate mismatch');
    });
  });

  group('decodeJson', () {
    final tai0 = Calends.decodeJson(
        '"40000000000000000000000000000000000000000000000000000000"');

    test('TAI64 date', () {
      expect(tai0.date('tai64', 'decimal'), equals('0'),
          reason: 'date mismatch');
    });

    test('duration', () {
      expect(tai0.duration(), equals('0'), reason: 'duration mismatch');
    });

    test('TAI64 endDate', () {
      expect(tai0.endDate('tai64', 'decimal'), equals('0'),
          reason: 'endDate mismatch');
    });
  });
}
