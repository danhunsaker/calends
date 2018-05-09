namespace Calends;

class TAITime implements \Serializable
{
  public seconds = 0.0;
  public nano = 0;
  public atto = 0;
  public xicto = 0;
  public ucto = 0;
  public rocto = 0;

  public function __construct(stamp = null)
  {
    var result;

    if empty stamp {
      let stamp = "";
    }

    switch typeof stamp {
      case "string":
        if substr(stamp, 0, 2) == "0x" {
          let result = TAI64Time_from_hex_string(stamp);
        } else {
          let result = TAI64Time_from_string(stamp);
        }
        break;
      case "integer":
      case "double":
        let result = TAI64Time_from_double((double)stamp);
        break;
      default:
        throw new CalendsException("Unsupported input type " . typeof stamp);
    }

    let this->seconds = result->seconds ?: 0.0;
    let this->nano = result->nano ?: 0;
    let this->atto = result->atto ?: 0;
    let this->xicto = result->xicto ?: 0;
    let this->ucto = result->ucto ?: 0;
    let this->rocto = result->rocto ?: 0;
  }

  public function __toString() -> string
  {
    return self::toString();
  }

  public function serialize() -> string
  {
    string out;

    let out = TAI64Time_encode_text(this);

    return out;
  }

  public function unserialize(data) -> void
  {
    var result;

    let result = TAI64Time_decode_text((string)data);

    let this->seconds = result->seconds;
    let this->nano = result->nano;
    let this->atto = result->atto;
    let this->xicto = result->xicto;
    let this->ucto = result->ucto;
    let this->rocto = result->rocto;
  }

  public function add(<TAITime> z) -> <TAITime>
  {
    var result;

    let result = TAI64Time_add(this, z);

    return result;
  }

  public function sub(<TAITime> z) -> <TAITime>
  {
    var result;

    let result = TAI64Time_sub(this, z);

    return result;
  }

  public function toString() -> string
  {
    string out;

    let out = TAI64Time_string(this);

    return out;
  }

  public static function fromString(string value = null) -> <TAITime>
  {
    var result;

    let result = TAI64Time_from_string(value);

    return result;
  }

  public function toHex() -> string
  {
    string out;

    let out = TAI64Time_hex_string(this);

    return out;
  }

  public static function fromHex(string value = null) -> <TAITime>
  {
    var result;

    let result = TAI64Time_from_hex_string(value);

    return result;
  }

  public function toNumber() -> double
  {
    var out;

    let out = TAI64Time_double(this);

    return (double)out;
  }

  public static function fromNumber(value = null) -> <TAITime>
  {
    var result;

    let result = TAI64Time_from_double((double)value);

    return result;
  }

  public function toUTC() -> <TAITime>
  {
    var result;

    let result = TAI64Time_tai_to_utc(this);

    return result;
  }

  public function fromUTC() -> <TAITime>
  {
    var result;

    let result = TAI64Time_utc_to_tai(this);

    return result;
  }
}
