# fsup

find the first file or path recursively upstream from a base path

# Warning

this does not stop until it finds a match or errors

this will find things WAY higher up the file tree than you may intend, so be careful of file operations outside of your intended scope

this library does not distinguish between folders, files, or any other type of stat... only if the os says it exists
