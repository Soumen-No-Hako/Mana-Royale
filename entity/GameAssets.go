package assets

type Player struct {
    HP int
    Name string
    DPT int
}

type Card struct {
    Name string
    Category string    // 3 categories - DMG, PROJECTILE, EPIC
    Value int
}
//add betrayal chance in the main code during initialization
