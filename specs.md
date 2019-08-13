# Presentation: 
https://prezi.com/view/Q5GV2R30SSrv4DOLUWn3/ 

# Application Specs

Proiectul pe care m-am gandit sa il dezvolt (se bazeaza  pe tutorialul din documentatia Cosmos) este o platforma prin care utilizatorii pot cumpara/vinde obiecte.
Fiecare utilizator al retelei are asociat un obiect, si un pret minim setat de catre detinator, impreuna reprezentand o structura.
Daca un alt utilizator, sa presupunem utilizatorul 2 va licita un pret mai mare sau egal cu cel minim, pentru obiectul detinut sa spunem de catre utilizatorul 1, noul proprietar al obiectului va fi utilizatorul 2.

Utilizatorii care vor sa doneze obiecte, nu vor fi trecuti ca si detinatori ai obiectelor, ci doar o structura a obiectului se va afla in retea, fara nici un detinator.
O data ce un utilizator al retelei doreste sa detina un obiect fara nume din retea, numele acestuia va fi trecut ca si detinator al obiectului, nefiind necesara nicio tranzactie. 
Se va seta doar noul nume in structura, si un pret minim.

Fiecare strutura va trece prin mai multe stari, iar tranzactiile vor fi cele care vor hotari schimbarea starilor.

Am incercat sa dezvolt acest proiect, bazandu-ma pe Cosmos SDK, din pacate nu am reusit sa finalizez, sau sa testez cel putin ceea ce am facut, de aceea nu am facut nici o prezentare cu rezultatele :).
Am reusit sa implementez o doar o stare, si anume starea initiala cand obiectele au sau nu detinator, dar nu cred ca e completa.
Urmatoarea stare ar fi fost MsgBuyObject(obiect, valoare, cumparator) si handler ul corespunzator.
Apoi ar mai fi fost partea de querier, parte pe care nu am prea intes cum functioneaza sau pentru ce este necesara.

## Implementare

Dupa ce am citit documentatia Cosmos SDK si temdermint, am inteles ca eu ca si developer ar trebui sa definesc doar starile strcuturilor din retea. Starea de inceput (genesis),
dar si mesajele care vor declansa tranzitiile spre alte stari, iar Tendermint se va ocupa de propagarea in retea. 

Am incercat sa las cateva comentarii prin cod, mai mult pentru mine pentru a ma ajuta in timp ce implementez, si pentru a face putina lumina in capul meu, deci s-ar putea sa fie putina dezordine.
