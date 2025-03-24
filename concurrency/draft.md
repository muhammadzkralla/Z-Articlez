# Project Java Cooperative Multitasking (Coroutines)

## Initial Functions:

### void yield();

Suspend the current coroutine without blocking the whole thread by switching the execution context to the next available coroutine ready to be exectued. It yields the thread (or thread pool) of the current coroutine dispatcher to other coroutines on the same dispatcher to run if possible. This requires saving the current runnable state in a state machine (continuation) to be restored back when the current coroutine resumes.
