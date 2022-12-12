import 'dart:ffi';
import 'dart:typed_data';

import 'package:ffi/ffi.dart';

import 'generated_bindings.dart' show TAI64Time;
import 'util.dart';

extension TAI64TimeMethods on TAI64Time {
  static TAI64Time fromDouble(double stamp) {
    return loadLib().TAI64Time_from_double(stamp);
  }

  static TAI64Time fromTAI64String(String stamp) {
    return loadLib().TAI64Time_from_string(strToCharPtr(stamp));
  }

  static TAI64Time fromHex(String stamp) {
    return loadLib().TAI64Time_from_hex_string(strToCharPtr(stamp));
  }

  static TAI64Time decodeText(String stamp) {
    return loadLib().TAI64Time_decode_text(strToCharPtr(stamp));
  }

  static TAI64Time decodeBinary(Uint8List stamp) {
    return loadLib()
        .TAI64Time_decode_binary(bytesToVoidPtr(stamp), stamp.lengthInBytes);
  }

  TAI64Time add(TAI64Time other) {
    return loadLib().TAI64Time_add(this, other);
  }

  TAI64Time sub(TAI64Time other) {
    return loadLib().TAI64Time_sub(this, other);
  }

  TAI64Time utcToTai() {
    return loadLib().TAI64Time_utc_to_tai(this);
  }

  TAI64Time taiToUtc() {
    return loadLib().TAI64Time_tai_to_utc(this);
  }

  double toDouble() {
    return loadLib().TAI64Time_double(this);
  }

  String toTAI64String() {
    return charPtrToStr(loadLib().TAI64Time_string(this));
  }

  String toHex() {
    return charPtrToStr(loadLib().TAI64Time_hex_string(this));
  }

  String encodeText() {
    return charPtrToStr(loadLib().TAI64Time_encode_text(this));
  }

  Uint8List encodeBinary() {
    final Pointer<Int> length = malloc<Int>(1);
    final data = loadLib().TAI64Time_encode_binary(this, length);
    return voidPtrToBytes(data, length.value);
  }
}
