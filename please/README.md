DOESN'T WORK

all the stuff in third_party is necessary - the proto_\* rules depend on it
But it's not built in
There's a version of it all in the main please repo, but you don't seem to be able to depend on http
So I copied that, but had to make a bunch of fixups
Their own tests in the main please repo don't work - hit the same issues I did
Even with my fixes, hit an error with protoc's build-in proto libs apparently not being available.
The please repo tests imply they'll just work, but other stuff in that repo is broken so
