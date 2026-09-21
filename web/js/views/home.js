import { createTeams } from "../createTeams.js";

export default function renderHome(app, state, refresh) {
    const selectedPlayerIds = new Set(
        state.teams.flatMap((match) => [
            ...(match.teamA || []).map((player) => player.id),
            ...(match.teamB || []).map((player) => player.id),
        ]),
    );
    const numberOfTeams = state.teams.length * 2 || 2;

    app.innerHTML = `
        <main class="content">
            <section class="controls">
                <label>
                    ¿Cuántos equipos sois?
                    <select id="team-count">
                        ${[2, 4, 6]
                            .map(
                                (value) => `
                            <option value="${value}" ${value === numberOfTeams ? "selected" : ""}>${value}</option>
                        `,
                            )
                            .join("")}
                    </select>
                </label>
                <span>Jugadores seleccionados: <strong id="selected-count">${selectedPlayerIds.size}</strong></span>
                <div class="button-group">
                    <button id="create-teams" type="button" ${selectedPlayerIds.size ? "" : "disabled"}>Crear equipos</button>
                    <button id="clear-teams" class="ghost" type="button">Borrar equipos</button>
                </div>
            </section>

            <section class="teams-container">${renderMatches(state.teams)}</section>

            <section class="table-section">
                <div class="desktop-view">
                    <div class="section-heading">
                        <h2>Jugadores</h2>
                        <button class="add-player" type="button">Añadir jugador</button>
                    </div>
                    ${renderTable(state.players, selectedPlayerIds)}
                </div>

                <div class="mobile-view">
                    <div class="section-heading">
                        <h2>Jugadores</h2>
                        <button class="add-player" type="button">Añadir</button>
                    </div>
                    <div class="cards-container">
                        ${state.players.map((player) => renderCard(player, selectedPlayerIds)).join("")}
                    </div>
                </div>
            </section>
        </main>
        <dialog id="player-dialog">
            <form id="player-form">
                <h2>Añadir jugador</h2>
                <label for="player-name">Nombre del jugador</label>
                <input id="player-name" required placeholder="Introduce el nombre del jugador...">
                <div class="dialog-actions">
                    <button class="ghost" type="button" id="cancel-player">Cancelar</button>
                    <button type="submit">Añadir jugador</button>
                </div>
            </form>
        </dialog>
    `;

    document
        .querySelector("#create-teams")
        .addEventListener("click", async () => {
            const selectedPlayers = state.players.filter((player) =>
                selectedPlayerIds.has(player.id),
            );
            state.teams = createTeams(
                selectedPlayers,
                Number(document.querySelector("#team-count").value),
            );
            await saveTeams(state.teams);
            renderHome(app, state, refresh);
        });

    document
        .querySelector("#clear-teams")
        .addEventListener("click", async () => {
            state.teams = [];
            await saveTeams([]);
            renderHome(app, state, refresh);
        });

    document.querySelectorAll("[data-player-id]").forEach((element) => {
        element.addEventListener("click", (event) => {
            if (event.target.closest("[data-delete-id]")) return;
            const id = Number(element.dataset.playerId);
            selectedPlayerIds.has(id)
                ? selectedPlayerIds.delete(id)
                : selectedPlayerIds.add(id);
            element.classList.toggle("selected", selectedPlayerIds.has(id));
            const checkbox = element.querySelector('input[type="checkbox"]');
            if (checkbox) checkbox.checked = selectedPlayerIds.has(id);
            document.querySelector("#selected-count").textContent =
                selectedPlayerIds.size;
            document.querySelector("#create-teams").disabled =
                selectedPlayerIds.size === 0;
        });
    });

    document.querySelectorAll("[data-delete-id]").forEach((button) => {
        button.addEventListener("click", async (event) => {
            event.stopPropagation();
            if (button.dataset.confirmed !== "true") {
                button.dataset.confirmed = "true";
                button.textContent = "✓";
                return;
            }
            await request("/api/players", {
                method: "DELETE",
                body: { id: Number(button.dataset.deleteId) },
            });
            await refresh();
        });
    });

    document.querySelectorAll("[data-score]").forEach((input) => {
        input.addEventListener("input", () => {
            const [index, side] = input.dataset.score.split(":");
            state.teams[index][side] =
                input.value === "" ? undefined : Number(input.value);
            input
                .closest(".team-match")
                .querySelector("[data-match-index]").disabled =
                state.teams[index].teamAScore === undefined ||
                state.teams[index].teamBScore === undefined;
        });
    });

    document.querySelectorAll("[data-match-index]").forEach((button) => {
        button.addEventListener("click", async () => {
            const match = state.teams[Number(button.dataset.matchIndex)];
            await request("/api/matches", {
                method: "POST",
                body: {
                    date: new Date(), // @TODO: Change it to the selected Date
                    seasonId: 1, // @TODO: Change it to the selected Season ID
                    teamA: match.teamA.map((player) => {
                        return { playerId: player.id };
                    }),
                    teamB: match.teamB.map((player) => {
                        return { playerId: player.id };
                    }),
                    teamAGoals: match.teamAScore,
                    teamBGoals: match.teamBScore,
                },
            });
            await refresh();
        });
    });

    const dialog = document.querySelector("#player-dialog");
    document
        .querySelectorAll(".add-player")
        .forEach((button) =>
            button.addEventListener("click", () => dialog.showModal()),
        );
    document
        .querySelector("#cancel-player")
        .addEventListener("click", () => dialog.close());
    document
        .querySelector("#player-form")
        .addEventListener("submit", async (event) => {
            event.preventDefault();
            await request("/api/players", {
                method: "POST",
                body: {
                    name: document.querySelector("#player-name").value.trim(),
                },
            });
            dialog.close();
            await refresh();
        });
}

function renderMatches(matches) {
    return matches
        .map(
            (match, index) => `
        <article class="team-match">
            <div class="teams-grid">
                ${renderTeam(`Equipo ${index * 2 + 1}`, match.teamA || [])}
                ${renderTeam(`Equipo ${index * 2 + 2}`, match.teamB || [])}
            </div>
            <div class="score-section">
                <label>Goles Equipo ${index * 2 + 1}<input type="number" min="0" data-score="${index}:teamAScore" value="${match.teamAScore ?? ""}"></label>
                <label>Goles Equipo ${index * 2 + 2}<input type="number" min="0" data-score="${index}:teamBScore" value="${match.teamBScore ?? ""}"></label>
                <button type="button" data-match-index="${index}" ${match.teamAScore === undefined || match.teamBScore === undefined ? "disabled" : ""}>Registrar</button>
            </div>
        </article>
    `,
        )
        .join("");
}

function renderTeam(title, players) {
    const winPercentage = players.length
        ? Math.trunc(
              players.reduce(
                  (sum, player) => sum + (player.victoryPercentage || 0),
                  0,
              ) / players.length,
          )
        : 0;
    return `<div class="team"><h3>${title} <small>(Jugadores: ${players.length})</small></h3><div class="players-list">${players.map((player) => `<p>${escapeHtml(player.name)}</p>`).join("")}</div><div class="stats">Victorias: <strong>${winPercentage}%</strong></div></div>`;
}

function renderTable(players, selected) {
    const headers = [
        "Nombre",
        "Ganados",
        "Empatados",
        "Perdidos",
        "Puntos",
        "Victoria %",
        "GF",
        "GC",
        "",
    ];
    return `<table><thead><tr>${headers.map((header) => `<th>${header}</th>`).join("")}</tr></thead><tbody>${players
        .map(
            (player) => `
        <tr class="${selected.has(player.id) ? "selected" : ""}" data-player-id="${player.id}">
            <td>${escapeHtml(player.name)}</td><td>${player.matchesWon ?? ""}</td><td>${player.matchesDrawn ?? ""}</td><td>${player.matchesLost ?? ""}</td><td>${player.totalPoints ?? ""}</td><td>${player.victoryPercentage ? `${Math.trunc(player.victoryPercentage)}%` : "0%"}</td><td>${player.goalsFor ?? ""}</td><td>${player.goalsAgainst ?? ""}</td><td><button class="delete" type="button" data-delete-id="${player.id}">×</button></td>
        </tr>`,
        )
        .join("")}</tbody></table>`;
}

function renderCard(player, selected) {
    return `<article class="player-card ${selected.has(player.id) ? "selected" : ""}" data-player-id="${player.id}">
        <div class="card-header"><div><h3>${escapeHtml(player.name)}</h3><p>${player.victoryPercentage ? Math.trunc(player.victoryPercentage) : 0}% victorias</p></div><button class="delete" type="button" data-delete-id="${player.id}">×</button></div>
        <div class="card-stats"><span>Ganados <b>${player.matchesWon ?? ""}</b></span><span>Empatados <b>${player.matchesDrawn ?? ""}</b></span><span>Perdidos <b>${player.matchesLost ?? ""}</b></span><span>Puntos <b>${player.totalPoints ?? ""}</b></span></div>
        <div class="card-footer"><span>⚽ ${player.goalsFor ?? ""} - ${player.goalsAgainst ?? ""}</span><input type="checkbox" ${selected.has(player.id) ? "checked" : ""} tabindex="-1" aria-label="Seleccionar jugador"></div>
    </article>`;
}

async function saveTeams(teams) {
    await request("/api/teams", { method: "POST", body: { teams } });
}

async function request(url, options = {}) {
    const response = await fetch(url, {
        ...options,
        headers: {
            "content-type": "application/json",
            ...(options.headers || {}),
        },
        body:
            options.body === undefined
                ? undefined
                : JSON.stringify(options.body),
    });
    if (!response.ok)
        throw new Error(
            `${options.method || "GET"} ${url} failed (${response.status})`,
        );
    return response;
}

function escapeHtml(value) {
    return String(value)
        .replaceAll("&", "&amp;")
        .replaceAll("<", "&lt;")
        .replaceAll(">", "&gt;")
        .replaceAll('"', "&quot;")
        .replaceAll("'", "&#039;");
}
