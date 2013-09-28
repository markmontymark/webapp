use JSON::XS;
use Task;

package HttpTask;
use Mo qw'build default builder coerce is required';
extends 'Task';

has method => ();
has url => ();
has query => ();

sub run
{
	my $self = shift;

	my $do_pre_echo = $self->method && $self->method =~ /POST/i;

	my $process;
	if($self->query)
	{
		my @retval = ();
		for(@{$self->query})
		{
			$process = $do_pre_echo ? 
				sprintf("echo '%s'|%s %s", JSON::XS::encode_json($_), $self->method , $self->url) :
				sprintf("%s %s", $self->method , $self->url);
			print STDERR "$process\n";
			push @retval, `$process`;
		}
		return join "\n",@retval;
	} 
	$process = $self->method . ' ' . $self->url;
	print STDERR "$process\n";
	return `$process`;
}

1;
