$(document).ready( function () {


    var audioElement = document.createElement('audio');
    audioElement.setAttribute('src', 'https://eventku.herokuapp.com/sound/khitan.mp3');
    
    audioElement.addEventListener('ended', function() {
        this.play();
    }, false);
    
    audioElement.addEventListener("canplay",function(){
        $("#length").text("Duration:" + audioElement.duration + " seconds");
        $("#source").text("Source:" + audioElement.src);
        $("#status").text("Status: Ready to play").css("color","green");
    });
    
    audioElement.addEventListener("timeupdate",function(){
        $("#currentTime").text("Current second:" + audioElement.currentTime);
    });
    audioElement.play();
    $('#play').click(function() {
        audioElement.play();
        $("#status").text("Status: Playing");
    });
    
    $('#pause').click(function() {
        audioElement.pause();
        $("#status").text("Status: Paused");
    });
    
    $('#restart').click(function() {
        audioElement.currentTime = 0;
    });

    // var audio = {};
    //     audio["walk"] = new Audio();
    //     audio["walk"].src = "/sound/khitan.mp3"
    //     audio["walk"].addEventListener('load', function () {
    //         audio["walk"].play();
    //     });

// var myaudio = $("#audioID");
// myaudio.play();
  } );