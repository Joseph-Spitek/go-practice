package main

func gradingStudents(grades []int32) []int32 {
    for i, grade := range grades {
        rem := grade % 5
        if (rem > 2 && grade + 5 - rem >= 40)  {
            grades[i] = grade + 5 - rem
        }
    }
    return grades
}
