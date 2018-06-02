/*
Package fsv aims to simplify file IO and interaction with the file system.
It revolves around the main type declared, which is fsv.Path and offers
a lot of methods, to easily retrieve information or do manipulations on the fs.
The names of methods are largely inspired by bash/Linux commands, such as cp,
mv or mkdir. Furthermore, some methods are slightly overloaded in the sense that
their behaviour can be modified by passing flags.
*/
package fsv
