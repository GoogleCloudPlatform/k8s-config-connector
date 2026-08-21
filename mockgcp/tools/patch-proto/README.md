This is a simple tool to manipulate .proto files.

We use it to patch in missing fields into the public googleapis protos,
until the public protos are updated.

It uses the treesitter library, which is arguably overkill for what
we're doing right now.  But we are increasingly manipulating and inspecting
proto and go files, so this seems like a good sandbox in which to evaluate
treesitter.