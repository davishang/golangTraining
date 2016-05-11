//swithcing images
var myImage = document.querySelector('img');

myImage.onclick = function() {
    var mySrc = myImage.getAttribute('src');
    if(mySrc === 'images/mural1.jpg') {
        myImage.setAttribute ('src','images/mural2.jpg');
    } else {
        myImage.setAttribute ('src','images/mural1.jpg');
    }
}

//personalizing message code
var myButton = document.querySelector('button');
var myHeading = document.querySelector('h1');

function setUserName() {
    var myName = prompt('Please enter your Superhero name.');
    localStorage.setItem('name', myName);
    myHeading.innerHTML = myName + ' was here.';
}

if(!localStorage.getItem('name')) {
    setUserName();
} else {
    var storedName = localStorage.getItem('name');
    myHeading.innerHTML = storedName + ' is cool.';
}

myButton.onclick = function() {
    setUserName();
};