package airportrobot

import "fmt"

// Greeter interface defines methods for language-specific greetings
type Greeter interface {
    LanguageName() string
    Greet(visitorName string) string
}

// Italian implements the Greeter interface for Italian language
type Italian struct{}

// Portuguese implements the Greeter interface for Portuguese language
type Portuguese struct{}

// LanguageName returns the name of the Italian language
func (i Italian) LanguageName() string {
    return "Italian"
}

// Greet returns an Italian greeting for the given visitor name
func (i Italian) Greet(visitorName string) string {
    const italianGreeting = "Ciao %s!"
    return fmt.Sprintf(italianGreeting, visitorName)
}

// LanguageName returns the name of the Portuguese language
func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

// Greet returns a Portuguese greeting for the given visitor name
func (p Portuguese) Greet(visitorName string) string {
    const portugueseGreeting = "Olá %s!"
    return fmt.Sprintf(portugueseGreeting, visitorName)
}

// SayHello creates a complete greeting message using the provided greeter
func SayHello(visitorName string, greeter Greeter) string {
    const greetingFormat = "I can speak %s: %s"
    return fmt.Sprintf(greetingFormat, greeter.LanguageName(), greeter.Greet(visitorName))
}
