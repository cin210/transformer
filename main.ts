
# Simply start by parsing the document as 'html'
html() {
  # Select the body
  $("/html/body") {
    # Append a class to the body tag
    add_class("moov")
    # Caturday
    $$("img"){
      attr("src", "http://thecatapi.com/api/images/get?format=src")
    }
    # Sometimes empty blocks break stuff
    $(".") {
    }
  }
}