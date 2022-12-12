import 'dart:typed_data';

import 'package:calends/calends.dart';
import 'package:test/test.dart';

void main() {
  group('TAI64TimeMethods 0 fromDouble', () {
    final tai0 = TAI64TimeMethods.fromDouble(0);

    test('Seconds', () {
      expect(tai0.Seconds, equals(0));
    });

    test('Nano', () {
      expect(tai0.Nano, equals(0));
    });

    test('Atto', () {
      expect(tai0.Atto, equals(0));
    });

    test('Xicto', () {
      expect(tai0.Xicto, equals(0));
    });

    test('Ucto', () {
      expect(tai0.Ucto, equals(0));
    });

    test('Rocto', () {
      expect(tai0.Rocto, equals(0));
    });

    test('toDouble', () {
      expect(tai0.toDouble(), equals(0));
    });
  });

  group('TAI64TimeMethods 0 fromTAI64String', () {
    final tai0 = TAI64TimeMethods.fromTAI64String('0');

    test('Seconds', () {
      expect(tai0.Seconds, equals(0));
    });

    test('Nano', () {
      expect(tai0.Nano, equals(0));
    });

    test('Atto', () {
      expect(tai0.Atto, equals(0));
    });

    test('Xicto', () {
      expect(tai0.Xicto, equals(0));
    });

    test('Ucto', () {
      expect(tai0.Ucto, equals(0));
    });

    test('Rocto', () {
      expect(tai0.Rocto, equals(0));
    });

    test('toTAI64String', () {
      expect(tai0.toTAI64String(), equals('0'));
    });
  });

  group('TAI64TimeMethods 0 fromHex', () {
    final tai0 = TAI64TimeMethods.fromHex(
        '40000000000000000000000000000000000000000000000000000000');

    test('Seconds', () {
      expect(tai0.Seconds, equals(0));
    });

    test('Nano', () {
      expect(tai0.Nano, equals(0));
    });

    test('Atto', () {
      expect(tai0.Atto, equals(0));
    });

    test('Xicto', () {
      expect(tai0.Xicto, equals(0));
    });

    test('Ucto', () {
      expect(tai0.Ucto, equals(0));
    });

    test('Rocto', () {
      expect(tai0.Rocto, equals(0));
    });

    test('toHex', () {
      expect(tai0.toHex(),
          equals('40000000000000000000000000000000000000000000000000000000'));
    });
  });

  group('TAI64TimeMethods 0 decodeText', () {
    final tai0 = TAI64TimeMethods.decodeText(
        '40000000000000000000000000000000000000000000000000000000');

    test('Seconds', () {
      expect(tai0.Seconds, equals(0));
    });

    test('Nano', () {
      expect(tai0.Nano, equals(0));
    });

    test('Atto', () {
      expect(tai0.Atto, equals(0));
    });

    test('Xicto', () {
      expect(tai0.Xicto, equals(0));
    });

    test('Ucto', () {
      expect(tai0.Ucto, equals(0));
    });

    test('Rocto', () {
      expect(tai0.Rocto, equals(0));
    });

    test('encodeText', () {
      expect(tai0.encodeText(),
          equals('40000000000000000000000000000000000000000000000000000000'));
    });
  });

  group('TAI64TimeMethods 0 decodeBinary', () {
    final tai0 = TAI64TimeMethods.decodeBinary(Uint8List.fromList([
      64,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      0
    ]));

    test('Seconds', () {
      expect(tai0.Seconds, equals(0));
    });

    test('Nano', () {
      expect(tai0.Nano, equals(0));
    });

    test('Atto', () {
      expect(tai0.Atto, equals(0));
    });

    test('Xicto', () {
      expect(tai0.Xicto, equals(0));
    });

    test('Ucto', () {
      expect(tai0.Ucto, equals(0));
    });

    test('Rocto', () {
      expect(tai0.Rocto, equals(0));
    });

    test('encodeBinary', () {
      expect(
          tai0.encodeBinary(),
          equals(Uint8List.fromList([
            64,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0,
            0
          ])));
    });
  });

  group('TAI64TimeMethods 0', () {
    final tai0 = TAI64TimeMethods.fromDouble(0);

    test('add', () {
      final tai1 = TAI64TimeMethods.fromDouble(1);

      expect(tai0.add(tai1).Seconds, equals(1));
    });

    test('sub', () {
      final tai1 = TAI64TimeMethods.fromDouble(1);

      expect(tai0.sub(tai1).Seconds, equals(-1));
    });

    test('taiToUtc', () {
      expect(tai0.taiToUtc().Seconds, equals(8));
    });

    test('utcToTai', () {
      expect(tai0.utcToTai().Seconds, equals(-7));
    });
  });
}
