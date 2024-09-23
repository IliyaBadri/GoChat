function CheckUsername(username){
    const regex = /^[A-Za-z0-9]{3,30}$/;

    return regex.test(username);
}

const usernameInput = document.getElementById("username_input");
const loginButton = document.getElementById("login_button");
const instructions = document.getElementById("instructions");

loginButton.addEventListener("click", () => {
    if(!CheckUsername(usernameInput.value)) {
        instructions.style.display = "block";
        ChangeStatus("Invalid username. Please read the instructions and try again.");
    } else {
        instructions.style.display = "none";
        usernameInput.value = "";
        ChangeStatus("Logging in . . .");
    }
        
});

ChangeStatus("Waiting for you to login . . .");