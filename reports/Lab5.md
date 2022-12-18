# Web Authentication and authorisation

## Course: Cryptography & Security

## Author: Zacatov Andrei

----

### Theory

Authorisation and authentication are 2 of the most important main security goals in IT systems.
They are closely related, yet one should know the difference between the two.
Authentication is the process of verifying the identity of a user, ansering the question of "who are you?".
Authorisation deals with checking the access rights, and mostly answers the question of "what can you do in the system"

### Objectives

1. Create a web service
2. Service should implement some basic authentication.
3. I could not implement MFA and i believe it would lose me a couple points, but i suppose it is better that way.
4. service needs to simulate user authorisation

### Implementation description

1. The authentication
    1. The authentication in my project will be simple and based on HTTP cookie.
       This is not the best way, especially for a non-HTTPS communication where data is at most Base32 encrypted.
       However it provides relative simplicity since i do not know any web frameworks.
    2. A middleware concept has been used in this example. Middlewares can be regarded as a "chain of responsibility"
       pattern in software design.
       The middlewares share a "common" signature, that being
        ```go
       func Middleware(next http.HandlerFunc /*specific params*/) http.HandlerFunc
         ```
       that would allow us to chain them one inside another.
    3. The authentication middleware code:
       ```go
        func Authenticate(next http.HandlerFunc, DB *passwords.UserDB) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            username, password, ok := r.BasicAuth()
            if ok && DB.Authenticate(username, password) {
                next.ServeHTTP(w, r)
            return
            }
        http.Error(w, "None/Incorrect auth data provided", http.StatusUnauthorized)
            }
        }

        ```
       It checks the provided "WWW-AUTHENTICATE" header for correct input. Then if checks if the user is present in the
       Database.
       In case one of the tests fails it send a 401 `UNAUTHORISED` response

2. The authorisation
    1. The Authorisation works in a similar Middleware way.
    2. Code:
       ```go
            func Authorise(next http.HandlerFunc, DB *passwords.UserDB, allowedRoles ...string) http.HandlerFunc {
            return func(w http.ResponseWriter, r *http.Request) {
                userName, _, _ := r.BasicAuth()
                userRole, err := DB.GetRole(userName)
                if err != nil {
                    for _, role := range allowedRoles {
                        if role == userRole {
                            next.ServeHTTP(w, r)
                            return
                        }
                    }
                }
                http.Error(w, "Insufficient permissions", http.StatusForbidden)
                }
            }
        ```
       Here we also add a allowedRoles parameter that is a variable length parameter.
       However it can also be used with go slices using the `...` slice unpacking.
       The role is gotten from the database and if it is in the allowedRoles then the request is passed along the chain.
       otherwise the user is given a `403 Forbidden` response

3. The service example
    1. Here the Atbash cipher will be used as a service.
       User will have to POST the data to the `/atbash` endpoint and in case the user has been authorised they will get
       a response with the encoded Message.
    2. For that we will have to define a handler in a following way
         ```go
        http.HandleFunc("/atbash",
            server.Authenticate(
                server.Authorise(
                    server.AtbashHandler,
                    DataBase,
                    "admin",
                ),
                DataBase,
            ),
        )
         ```
       Where we first call the Authenticate middleware,then the authorise middleware, and only then the page handler,
       which looks the following
        ```go
         func AtbashHandler(w http.ResponseWriter, r *http.Request) {
             plainByte, err := io.ReadAll(r.Body)
             plaintext := string(plainByte)
             atbash := classical.NewAtbash("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
             fmt.Fprintln(w, atbash.Encode(plaintext))
         return
         }
         ```
       In this case the following page is only accessible by admins, however one can easily grant the access to a
       regular user by chaging the server.Authorise call to
        ```go
        http.HandleFunc("/atbash",
            server.Authenticate(
                server.Authorise(
                    server.AtbashHandler,
                    DataBase,
                    "admin",
                    "user",
                ),
                DataBase,
            ),
        )
       // or, perhaps
       allowedUsers := []string{"user","admin"}
        http.HandleFunc("/atbash",
            server.Authenticate(
                server.Authorise(
                    server.AtbashHandler,
                    DataBase,
                    allowedUsers...,
                ),
                DataBase,
            ),
        )

        ```

### Conclusions / Screenshots / Results

By the end of this laboratory work one has researched and implemented the ways of authenticating and authorising the
users in their own web app and has familiarised with the ways of implementing it.
At the end of it, one has tackled the MFA as a better,stronger and a more secure solution.
