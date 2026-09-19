function Router() {
    this.routes = {};
    this.currentRoute = "";

    this.addRoute = function (path, view) {
        this.routes[path] = view;
    };

    this.navigate = function (path) {
        if (this.routes[path]) {
            this.currentRoute = path;
            this.renderView(this.routes[path]);
            window.history.pushState({}, "", path);
        } else {
            console.error("Route not found: " + path);
        }
    };

    this.renderView = function (view) {
        const app = document.getElementById("app");
        app.innerHTML = ""; // Clear the current view
        app.appendChild(view()); // Render the new view
    };

    window.onpopstate = () => {
        this.navigate(window.location.pathname);
    };
}

const router = new Router();

// Example of adding a route
// router.addRoute('/', homeView); // homeView should be a function that returns the home view element

export default router;
