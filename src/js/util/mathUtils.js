define(function() {
  function toRadian(angleInDegrees) {
    return (Math.PI * angleInDegrees) / 180
  }

  //export
  return {
    toRadian: toRadian,
  }
});
