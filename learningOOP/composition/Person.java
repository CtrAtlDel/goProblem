package composition;

public class Person {
    private String name;

    public String getName() {
        return this.name;
    }
}

class Saiyan {
    private Person person;

    public String getName() {
        return this.person.getName();
    }
}