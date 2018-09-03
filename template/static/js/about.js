function onClickMe() {
  var xhr = new XMLHttpRequest()
  xhr.open("POST", '/api/v1/say/hello', true)
  
  xhr.onreadystatechange = function() {
    if (xhr.readyState === 4) {
      alert('Server: ' + this.response)
    }
  }
  xhr.send()
}
