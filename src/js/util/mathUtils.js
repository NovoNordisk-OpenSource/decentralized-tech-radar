define(function() {
  function toRadian(angleInDegrees) {
    return (Math.PI * angleInDegrees) / 180
  }

  return {
    toRadian: toRadian,
  }
});
