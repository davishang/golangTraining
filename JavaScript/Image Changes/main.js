var myImage = document.querySelector('img');

myImage.onclick = function() {
    var mySrc = myImage.getAttribute('src');
    if(mySrc === 'images/mural1.jpg') {
        myImage.setAttribute ('src','images/mural2.jpg');
    } else {
        myImage.setAttribute ('src','images/mural1.jpg');
    }
}