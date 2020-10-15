//Author: Prakhar Shukla


let News_Frame = document.getElementById('Present_frame');
let Next_News_Frame = document.getElementById('nextlink');
let Previous_News_Frame = document.getElementById('previouslink');




//Detecting Swipe Gesture
let pageWidth = window.innerWidth || document.body.clientWidth;
let threshold = Math.max(1, Math.floor(0.01 * (pageWidth)));
let touch_StartX = 0;
let touch_startY = 0;
let touch_endX = 0;
let touch_endY = 0;




const limit = Math.tan(45 * 1.5 / 180 * Math.PI);
const gestureZone = document.getElementById('Present_frame');



gestureZone.addEventListener('touchstart', function(event) {
    touch_StartX = event.changedTouches[0].screenX;
    touch_startY = event.changedTouches[0].screenY;
}, false);



gestureZone.addEventListener('touchend', function(event) {
    touch_endX = event.changedTouches[0].screenX;
    touch_endY = event.changedTouches[0].screenY;
    handlegesture(event);
}, false);




function handlegesture(e) {
    let x = touch_endX - touch_StartX;
    let y = touch_endY - touch_startY;

    let xy = Math.abs(x / y);
    let yx = Math.abs(y / x);

    if (Math.abs(x) > threshold || Math.abs(y) > threshold) {
        if (yx <= limit) {
            if (x < 0) {
                console.log("Left");
            } else {
                console.log("Right");
            }
        }
        if (xy <= limit) {
            if (y < 0) {
                console.log("top");
                if (Next_News_Frame !== "") {
                    window.location.replace(Next_News_Frame.href);
                }
            } else {
                console.log("bottom");
                if (Previous_News_Frame !== "") {
                    window.location.replace(Previous_News_Frame.href);
                }
            }
        }
    } else {
        console.log("tap");
    }
}