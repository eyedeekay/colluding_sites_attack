var hasConsole = typeof console !== "undefined"

var fingerprintReport = function () {
  var d1 = new Date()
  Fingerprint2.get(function(components) {
    var murmur = Fingerprint2.x64hash128(components.map(function (pair) { return pair.value }).join(), 31)
    var d2 = new Date()
    var time = d2 - d1
    document.querySelector("#time").textContent = time
    document.querySelector("#fp").textContent = murmur
    var details = ""
    if(hasConsole) {
      console.log("time", time)
      console.log("fingerprint hash", murmur)
    }
    for (var index in components) {
      var obj = components[index]
      var line = obj.key + " = " + String(obj.value).substr(0, 100)
      if (hasConsole) {
        console.log(line)
      }
      details += line + "\n"
    }
    document.querySelector("#details").textContent = details
  })
}

var cancelId
var cancelFunction

// see usage note in the README
if (window.requestIdleCallback) {
  cancelId = requestIdleCallback(fingerprintReport)
  cancelFunction = cancelIdleCallback
} else {
  cancelId = setTimeout(fingerprintReport, 500)
  cancelFunction = clearTimeout
}

document.querySelector("#btn").addEventListener("click", function () {
  if (cancelId) {
    cancelFunction(cancelId)
    cancelId = undefined
  }
  fingerprintReport()
})
