(() => {
    window.addEventListener('load', (event) => {
        const btnLogin = document.getElementById('btn-login'); 
        btnLogin.addEventListener('click', onLogin);
    });

    async function onLogin(event) {
        event.preventDefault();
        window.location.href = 'http://localhost:8080/oauth';
    }
})();
