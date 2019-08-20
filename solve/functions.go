package solve




func Topostfix(expr string) string{

	var s=new (Stack);
	var ans string="";
	for i,_:=range expr{

		if expr[i]=='1'||expr[i]=='0' {
			ans+=string(expr[i]);
		}else if expr[i]=='('{
			s.Push(string((expr[i])));
		} else if expr[i]==')'{
			for ; s.top.value!="(";{
				ans+=s.top.value;
				s.Pop()
			}
			s.Pop();
  		}else if expr[i]=='o'{
  			for; s.size>0&&(s.top.value=="o"||s.top.value=="a");{
  				ans+=s.top.value;s.Pop();
			}
  			s.Push("o");i++;
		}else if expr[i]=='a'{
			for; s.size>0&&(s.top.value=="o"||s.top.value=="a");{
				ans+=s.top.value;s.Pop();
			}
			s.Push("a");i+=2;
		}

	}
	for;s.size>0;{
		ans+=s.top.value;s.Pop();
	}

	return ans;
}
func Eva(expr string) string{
	var s=new(Stack);
	for i,_:=range expr{
		if expr[i]=='1'||expr[i]=='0'{
			s.Push(string(expr[i]));
		} else {
			a:=s.top.value;
			s.Pop();
			b:=s.top.value;
			s.Pop();
			 if expr[i]=='o' {
					 s.Push(string(Or(a, b)));
				 }
				 if expr[i] == 'a' {
					 s.Push(string(And(a, b)));
				 }
			 }
	}
	return s.top.value
}

func Or(a ,b string)string{

	if a=="1"||b=="1"{return "1";}
	return "0";
}
func And(a ,b string)string{

	if a=="1"&&b=="1"{return "1";}
	return "0";
}
