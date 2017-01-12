var variable = null;

for (var i = 0; i < 10; i++) {
    if (variable == null) {
        break;
    }
    variable = hazEsto(variable);
    if (variable == null) {
        break;
    }
    yEsto(variable);
    yLoOtro(variable);
    yMasCosas(variable);
}

for (var i = 0; i < 10; i++) {
    if (variable != null) {
        variable = hazEsto(variable);
        if (variable != null) {
            yEsto(variable);
            yLoOtro(variable);
            yMasCosas(variable);
        }
    }
}