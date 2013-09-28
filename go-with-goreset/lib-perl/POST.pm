use HttpTask;

package POST;
use Mo qw'build default builder coerce is required';
extends 'HttpTask';
has method => (default => sub{'POST'});

1;
