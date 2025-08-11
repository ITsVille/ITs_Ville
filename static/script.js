// Existing sayHello function
function sayHello() {
    alert("Hei! Voit laittaa viestiä: ville@example.com");

}

// Add chat widget controls
let chatVisible = false;

function toggleChat() {
    if (chatVisible) {
        window.botpressWebChat.hideChat();
        chatVisible = false;
    } else {
        window.botpressWebChat.showChat();
        chatVisible = true;
    }
}

// Initialize chat in hidden state
document.addEventListener('DOMContentLoaded', function() {
    window.botpressWebChat.hideChat();
});