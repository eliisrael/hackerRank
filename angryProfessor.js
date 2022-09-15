// Hacker Rank problem: https://www.hackerrank.com/challenges/angry-professor/problem?isFullScreen=true
const k = 3
const a = [-2, -1, 0, 1, 2]

function angryProfessor(k, a) {

    let attendanceStudents = 0

    for (let index = 0; index < a.length; index++) {
        if (a[index] <= 0) attendanceStudents++
    }

    if (attendanceStudents >= k) return "NO"
    else return "YES"


}


console.log(angryProfessor(k, a))