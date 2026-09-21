import renderHome from "./views/home.js";

const app = document.getElementById("app");

const savedTheme = localStorage.getItem("theme");
if (savedTheme === "light") document.documentElement.classList.add("light-theme");

document.getElementById("theme-toggle")?.addEventListener("click", () => {
    document.documentElement.classList.toggle("light-theme");
    localStorage.setItem("theme", document.documentElement.classList.contains("light-theme") ? "light" : "dark");
});

const state = {
    players: [],
    teams: [],
    currentSeason: null,
};

async function request(url, options = {}) {
    const response = await fetch(url, {
        ...options,
        headers: {
            "content-type": "application/json",
            ...(options.headers || {}),
        },
        body: options.body === undefined ? undefined : JSON.stringify(options.body),
    });

    if (!response.ok) {
        throw new Error(`${options.method || "GET"} ${url} failed (${response.status})`);
    }

    return response;
}

async function loadData() {
    const seasonResponse = await fetch("/api/seasons/current");
    if (!seasonResponse.ok) {
        state.currentSeason = null;
        state.players = [];
        state.teams = [];
        return;
    }

    const [playersResponse, teamsResponse] = await Promise.all([
        request("/api/players"),
        fetch("/api/teams"),
    ]);

    state.players = (await playersResponse.json()).data;
    const teamsData = teamsResponse.ok ? (await teamsResponse.json()).data : [];
    state.teams = Array.isArray(teamsData) ? teamsData : teamsData.teams || [];
    state.currentSeason = (await seasonResponse.json()).data;
}

async function loadView(view = "home") {
    if (view !== "home") {
        app.innerHTML = "<h1>404 - Page Not Found</h1>";
        return;
    }

    app.innerHTML = "<p>Loading...</p>";

    try {
        await loadData();
        renderHome(app, state, loadView);
    } catch (error) {
        console.error(error);
        app.innerHTML = '<p class="error">No se pudieron cargar los datos.</p>';
    }
}

window.addEventListener("popstate", () => {
    loadView(location.pathname.split("/").filter(Boolean).pop() || "home");
});

document.addEventListener("click", (event) => {
    const link = event.target.closest("[data-route]");
    if (!link) return;

    event.preventDefault();
    history.pushState({}, "", link.href);
    loadView("home");
});

loadView();
