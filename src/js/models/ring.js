define(function() {
  const Ring = function (name, order) {
    var self = {}

    self.name = function () {
      return name
    }

    self.order = function () {
      return order
    }

    return self
  }
  
  return Ring
})
