use HttpTask;

package GET;
use Mo qw'build default builder coerce is required';
extends 'HttpTask';
has method => (default => sub{'GET'});

1;
