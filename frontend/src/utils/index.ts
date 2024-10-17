// regex for email validation
const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
export const validateEmail = (email: string) => {
    if (!emailRegex.test(email)) {
        alert('Please enter a valid email.');
        return false;
    }
    return true;
}

//regex for password validation, at least 8 characters in total, at least one special character, no restrictions for uppercase or lowercase or underscore
const passwordRegex = /^(?=.*[!@#$%_^&*])[a-zA-Z0-9!@#$%_^&*]{8,}$/;
export const validatePassword = (password: string) => {
    if (!passwordRegex.test(password)) {
        alert('Please enter a valid password, password must contain at least 8 characters and one special character.');
        return false;
    }
    return true;
}