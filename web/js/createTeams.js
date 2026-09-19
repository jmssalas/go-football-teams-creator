function balanceTwoTeams(players) {
    let bestDivision = { teamA: [], teamB: [] };
    let minDifference = Infinity;

    function backtrack(index, teamA, teamB, victoryA, victoryB) {
        if (Math.abs(teamA.length - teamB.length) > 1) return;

        if (index >= players.length) {
            const percentageA = victoryA / teamA.length || 0;
            const percentageB = victoryB / teamB.length || 0;
            const difference = Math.abs(percentageA - percentageB);

            if (difference < minDifference) {
                minDifference = difference;
                bestDivision = { teamA: [...teamA], teamB: [...teamB] };
            }
            return;
        }

        const player = players[index];
        const victoryPercentage = player.victoryPercentage || 0;

        teamA.push(player);
        backtrack(index + 1, teamA, teamB, victoryA + victoryPercentage, victoryB);
        teamA.pop();

        teamB.push(player);
        backtrack(index + 1, teamA, teamB, victoryA, victoryB + victoryPercentage);
        teamB.pop();
    }

    backtrack(0, [], [], 0, 0);
    return bestDivision;
}

function balanceThreeTeams(players) {
    let bestDivision = { team1: [], team2: [], team3: [] };
    let minDifference = Infinity;

    function backtrack(index, team1, team2, team3, victory1, victory2, victory3) {
        if (
            Math.abs(team1.length - team2.length) > 1 ||
            Math.abs(team2.length - team3.length) > 1 ||
            Math.abs(team1.length - team3.length) > 1
        )
            return;

        if (index >= players.length) {
            const difference = Math.max(victory1, victory2, victory3) - Math.min(victory1, victory2, victory3);

            if (difference < minDifference) {
                minDifference = difference;
                bestDivision = {
                    team1: [...team1],
                    team2: [...team2],
                    team3: [...team3],
                };
            }
            return;
        }

        const player = players[index];
        const victoryPercentage = player.victoryPercentage || 0;

        team1.push(player);
        backtrack(index + 1, team1, team2, team3, victory1 + victoryPercentage, victory2, victory3);
        team1.pop();

        team2.push(player);
        backtrack(index + 1, team1, team2, team3, victory1, victory2 + victoryPercentage, victory3);
        team2.pop();

        team3.push(player);
        backtrack(index + 1, team1, team2, team3, victory1, victory2, victory3 + victoryPercentage);
        team3.pop();
    }

    backtrack(0, [], [], [], 0, 0, 0);
    return bestDivision;
}

function shuffle(players) {
    const shuffled = [...players];
    for (let index = shuffled.length - 1; index > 0; index -= 1) {
        const randomIndex = Math.floor(Math.random() * (index + 1));
        [shuffled[index], shuffled[randomIndex]] = [shuffled[randomIndex], shuffled[index]];
    }
    return shuffled;
}

export function createTeams(players, numberOfTeams) {
    const shuffled = shuffle(players);
    const teams = [];

    if (numberOfTeams === 2) {
        const division = balanceTwoTeams(shuffled);
        teams.push(division);
    } else if (numberOfTeams === 4) {
        const groups = balanceTwoTeams(shuffled);
        const firstPair = balanceTwoTeams(groups.teamA);
        const secondPair = balanceTwoTeams(groups.teamB);
        teams.push(firstPair, secondPair);
    } else if (numberOfTeams === 6) {
        const groups = balanceThreeTeams(shuffled);
        teams.push(balanceTwoTeams(groups.team1), balanceTwoTeams(groups.team2), balanceTwoTeams(groups.team3));
    }

    return teams;
}
