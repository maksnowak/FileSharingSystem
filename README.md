# System udostępniania plików - projekt na przedmiot *Projektowanie i Integracja Systemów*

[![wakatime](https://wakatime.com/badge/github/maksnowak/FileSharingSystem.svg)](https://wakatime.com/badge/github/maksnowak/FileSharingSystem)

## Opis projektu

- System udostępniania plików umożliwiający użytkownikom udostępnianie i pobieranie plików, pozwalający wyszukiwać
  pliki (do których mamy dostęp) po nazwie, typie, rozmiarze oraz datach (modyfikacji, utworzenia).
- System będzie umożliwiał zabezpieczenie dostępu do folderów oraz pojedynczych plików hasłem. W komponencie bazy danych
  będą trzymane ścieżki do plików na komputerze, na którym osadzony jest program.
- System będzie można wyeksportować/zaimportować, aby umożliwić posadowienie programu na innym serwerze bez utraty
  zaszyfrowanych plików, kategorii i tagów.
- System będzie umożliwiał tworzenia kont użytkowników, gdzie każdy użytkownik będzie posiadał pliki prywatne,
  ogólnodostępne pliki publiczne oraz będzie mógł udostępnić swoje pliki innemu użytkownikowi. Będzie także prowadzona
  ewidencja dostępu do własnych plików użytkownika (ile było pobrań, kto i kiedy pobrał).

## Jak zacząć lokalnie?

> Należy wpierw zainstalować make, technologie frontend/backend oraz Docker

Należy sklonować repozytorium

```shell
git clone https://github.com/maksnowak/FileSharingSystem
cd FileSharingSystem
```

W środowisku Linux wystarczy wywołać polecenie

```shell
make
```

### Inny system?

Na dowolnym systemie wspierającym konteneryzację Docker można uruchomić cały stos aplikacyjny, używając jednej komendy:

```shell
docker compose up --build
```

## Wymagania funkcjonalne:
**Wymagania obowiązkowe:**
- Użytkownik będzie w stanie utworzyć własne konto
- Użytkownik będzie w stanie zalogować się na swoje konto, zmienić hasło i je usunąć
- Użytkownik będzie w stanie wysyłać i pobierać pliki
- Użytkownik będzie w stanie udostępniać pliki innemu użytkownikowi
- Użytkownik będzie w stanie tworzyć pliki prywatne, które są niewidoczne dla innych użytkowników
- Użytkownik będzie w stanie kategoryzować pliki oraz nadawać im tagi
- Użytkownik będzie w stanie szyfrować pliki hasłem
- Użytkownik będzie w stanie wyszukiwać pliki w bazie (do których ma dostęp)
    - po nazwie
    - po typie (.pdf, .zip, .txt)
    - po rozmiarze
    - po dacie modyfikacji
    - po dacie utworzenia
- Użytkownik będzie w stanie wykonywać podstawowe operacje na plikach, takie jak:
    - usunięcie pliku
    - zmienianie nazwy

**Wymagania opcjonalne:**
- Użytkownik będzie w stanie zobaczyć własną ewidencję (ślad cyfrowego aktywności tzn. ile było pobrań, kiedy pobrał)
- Użytkownik będzie w stanie wyszukiwać pliki w bazie (do których ma dostęp) po zawartości (w przypadku plików tekstowych)
- Użytkownik będzie w stanie tworzyć foldery w strukturze drzewiastej i umieszczać tam pliki
- Użytkownik będzie w stanie szyfrować foldery hasłem
- Użytkownik będzie w stanie wyeksportować stan swojego dysku (stan obejmuje strukturę dysku, zaszyfrowane pliki, tagi
  oraz kategorie)
- Użytkownik będzie w stanie zaimportować stan dysku z pliku
- Użytkownik będzie w stanie zmienić miejsce przechowywania pliku na dysku

## Use Cases:

- Utworzenie konta, żeby móc korzystać z platformy
- Wysyłanie plików
- Pobieranie plików
- Udostępnienie plików innemu użytkownikowi
- Zmiana nazwy pliku
- Usunięcie pliku
- Szyfrowanie folderów/plików
- Wyszukiwanie plików po metadanych lub zawartości
- Kategoryzacja/tagowanie plików
- Podejrzenie własnej ewidencji
- Wyeksportowanie/Zaimportowanie stanu dysku (np. żeby nie stracić danych podczas zmiany serwera)

## Architektura:

Projekt zostanie zaimplementowany w następującej architekturze mikroserwisowej:

- Client (frontend) + Serwer
- Mikroserwis 1: Przesyłanie plików
- Mikroserwis 2: Szyfrowanie plików
- Mikroserwis 3: Tworzenie kont

W razie potrzeby można zmniejszyć/zwiększyć ilość mikroserwisów.

## Technologie:

- Frontend: Vue.js 3.5
- Backend: Go 1.23
- Baza danych: MongoDB 7.0
- Konteneryzacja: Docker 27.3

### Środowisko Developerskie:

- VCS + wrzucanie releasów: GitHub
- CI/CD: Azure DevOps
- Repozytorium mavenowe: nie dotyczy
- Budowanie projektu: Makefile
    - Możliwość uruchomienia testów
    - Możliwość *zabicia* aplikacji
- Wspierane systemu: Ubuntu Linux
- IDE: Goland (backend), VsCode (frontend)
- Zarządzanie zadaniami: GitHub Issues
- Pokrycie kodu testami: `go test –cover`, Vitest

## Kontrybucje

<a href="https://github.com/maksnowak/FileSharingSystem/graphs/contributors">
<img src="https://contrib.rocks/image?repo=maksnowak/FileSharingSystem" />
</a>

Made with [contrib.rocks](https://contrib.rocks).